package tailmon

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/liov/hoper/go/v2/initialize"
	"github.com/liov/hoper/go/v2/protobuf/utils/errorcode"
	"github.com/liov/hoper/go/v2/utils/log"
	httpi "github.com/liov/hoper/go/v2/utils/net/http"
	gin_build "github.com/liov/hoper/go/v2/utils/net/http/gin"
	"github.com/liov/hoper/go/v2/utils/net/http/grpc/gateway"
	"github.com/liov/hoper/go/v2/utils/strings"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	gtrace "golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (s *Server) httpHandler() http.HandlerFunc {
	//默认使用gin
	ginServer := gin_build.Http(initialize.InitConfig.ConfUrl, "../protobuf/api/", s.GinHandle)

	if s.GraphqlResolve != nil {
		graphqlServer := handler.NewDefaultServer(s.GraphqlResolve)
		ginServer.Handle(http.MethodPost, "/api/graphql", func(ctx *gin.Context) {
			graphqlServer.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}
	var gatewayServer http.Handler
	if s.GatewayRegistr != nil {
		gatewayServer = gateway.Gateway(s.GatewayRegistr)
		/*	ginServer.NoRoute(func(ctx *gin.Context) {
			gatewayServer.ServeHTTP(
				(*httpi.ResponseRecorder)(unsafe.Pointer(uintptr(*(*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(ctx))+8))))),
				ctx.Request)
			ctx.Writer.WriteHeader(http.StatusOK)
		})*/
	}
	var excludes = []string{"/debug", "/api-doc"}
	return func(w http.ResponseWriter, r *http.Request) {
		for _, exclude := range excludes {
			if strings.HasPrefix(r.RequestURI, exclude) {
				ginServer.ServeHTTP(w, r)
				return
			}
		}

		now := time.Now()
		recorder := httpi.NewRecorder(w.Header())

		body, _ := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewReader(body))

		ginServer.ServeHTTP(recorder, r)
		if recorder.Code == http.StatusNotFound && gatewayServer != nil {
			recorder.Reset()
			gatewayServer.ServeHTTP(recorder, r)
		}

		// 提取 recorder 中记录的状态码，写入到 ResponseWriter 中
		w.WriteHeader(recorder.Code)
		if recorder.Body != nil {
			// 将 recorder 记录的 Response Body 写入到 ResponseWriter 中，客户端收到响应报文体
			w.Write(recorder.Body.Bytes())
		}

		accessLog(r.RequestURI, stringsi.ToString(body), stringsi.ToString(recorder.Body.Bytes()),
			r.Header.Get(httpi.HeaderAuthorization), now, recorder.Code)
	}
}

type CustomContext func(c context.Context, r *http.Request) context.Context

func (s *Server) Serve(customContext CustomContext) {
	//反射从配置中取port
	serviceConfig := initialize.InitConfig.GetServiceConfig()

	if s.GRPCServer != nil {
		reflection.Register(s.GRPCServer)
	}
	httpHandler := s.httpHandler()
	openTracing := serviceConfig.OpenTracing
	systemTracing := serviceConfig.SystemTracing
	handle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.CallTwo.With(zap.String("stack", stringsi.ToString(debug.Stack()))).Error(" panic: ", r)
				w.Write(errorcode.SysErr)
			}
		}()

		// 请求TraceID，链路跟踪用
		ctx := r.Context()
		if systemTracing {
			// 系统trace只能追踪单个请求，且只记录时间及是否完成
			t := gtrace.New(initialize.InitConfig.Module, r.RequestURI)
			defer t.Finish()
			ctx = gtrace.NewContext(ctx, t)
		}
		if openTracing {
			var span *trace.Span
			// 直接从远程读取Trace信息，Trace是否为空交给propagation包判断
			if parent, ok := propagation.FromBinary(stringsi.ToBytes(r.Header.Get(httpi.HeaderTrace))); ok {
				ctx, span = trace.StartSpanWithRemoteParent(ctx, r.RequestURI,
					parent, trace.WithSampler(trace.AlwaysSample()),
					trace.WithSpanKind(trace.SpanKindServer))
			} else {
				ctx, span = trace.StartSpan(ctx, r.RequestURI,
					trace.WithSampler(trace.AlwaysSample()),
					trace.WithSpanKind(trace.SpanKindServer))
			}
			defer span.End()
		}

		if customContext != nil {
			r = r.WithContext(customContext(ctx, r))
		} else {
			r = r.WithContext(ctx)
		}

		if r.ProtoMajor == 2 && s.GRPCServer != nil && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			s.GRPCServer.ServeHTTP(w, r) // gRPC Server
		} else {
			httpHandler(w, r)
		}
	})
	h2Handler := h2c.NewHandler(handle, new(http2.Server))
	server := &http.Server{
		Addr:         serviceConfig.Port,
		Handler:      h2Handler,
		ReadTimeout:  serviceConfig.ReadTimeout,
		WriteTimeout: serviceConfig.WriteTimeout,
	}
	// 服务注册
	initialize.InitConfig.Register()
	//服务关闭
	cs := func() {
		if s.GRPCServer != nil {
			s.GRPCServer.Stop()
		}
		if err := server.Close(); err != nil {
			log.Error(err)
		}
	}
	go func() {
		<-close
		log.Debug("关闭服务")
		cs()
		close <- syscall.SIGINT
	}()

	go func() {
		<-stop
		log.Debug("重启服务")
		cs()
	}()
	log.Debugf("listening%v", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	GRPCServer     *grpc.Server
	GatewayRegistr gateway.GatewayHandle
	GinHandle      func(engine *gin.Engine)
	GraphqlResolve graphql.ExecutableSchema
}

var close = make(chan os.Signal, 1)
var stop = make(chan struct{}, 1)

func (s *Server) Start(newContext CustomContext) {
	if initialize.InitConfig.ConfigCenter == nil {
		log.Fatal(`初始化配置失败:
	main 函数的第一行应为
	defer v2.Start(config.Conf, dao.Dao)()
`)
	}
	signal.Notify(close,
		// kill -SIGINT XXXX 或 Ctrl+c
		syscall.SIGINT, // register that too, it should be ok
		// os.Kill等同于syscall.Kill
		syscall.SIGKILL, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
Loop:
	for {
		select {
		case <-close:
			break Loop
		default:
			s.Serve(newContext)
		}
	}
}

func ReStart() {
	stop <- struct{}{}
}

func accessLog(iface, body, result, auth string, start time.Time, code int) {
	log.Default.Logger.Info("", zap.String("interface", iface),
		zap.String("body", body),
		zap.Duration("processTime", time.Now().Sub(start)),
		zap.String("result", result),
		zap.String("auth", auth),
		zap.Int("status", code),
		zap.String("source", initialize.InitConfig.Module))
}
