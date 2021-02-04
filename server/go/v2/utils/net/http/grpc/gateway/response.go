package gateway

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/liov/hoper/go/v2/protobuf/utils/errorcode"
	"github.com/liov/hoper/go/v2/protobuf/utils/response"
	"github.com/liov/hoper/go/v2/utils/encoding/protobuf/jsonpb"
	httpi "github.com/liov/hoper/go/v2/utils/net/http"
	"github.com/liov/hoper/go/v2/utils/net/http/grpc/reconn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func ResponseHook(ctx context.Context, writer http.ResponseWriter, message proto.Message) error {
	if res, ok := message.(*response.HttpResponse); ok {
		for k, v := range res.Header {
			writer.Header().Add(k, v)
		}
		writer.WriteHeader(int(res.StatusCode))
	}
	/*	if message == nil{
		*(&message) = &response.TinyRep{Message: "OK"}
	}*/
	return nil
}

func HTTPError(ctx *gin.Context, err error) {

	s, ok := status.FromError(err)
	if ok && s.Code() == 14 && strings.HasSuffix(s.Message(), `refused it."`) {
		//提供一个思路，这里应该是哪条连接失败重连哪条，不能这么粗暴，map的key是个关键
		if len(reconn.ReConnectMap) > 0 {
			for _, f := range reconn.ReConnectMap {
				f()
			}
		}
	}

	const fallback = `{"code": 14, "message": "failed to marshal error message"}`

	delete(ctx.Request.Header,httpi.HeaderTrailer)
	contentType := jsonpb.JsonPb.ContentType(nil)
	ctx.Header(httpi.HeaderContentType, contentType)
	se, ok := err.(*errorcode.ErrRep)
	if !ok {
		se = &errorcode.ErrRep{Code: errorcode.Unknown, Message: err.Error()}
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	handleForwardResponseServerMetadata(ctx.Writer, md)

	buf, merr := jsonpb.JsonPb.Marshal(se)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", se, merr)
		ctx.Status(http.StatusInternalServerError)
		if _, err := io.WriteString(ctx.Writer, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	var wantsTrailers bool

	if te := ctx.GetHeader("TE"); strings.Contains(strings.ToLower(te), "trailers") {
		wantsTrailers = true
		handleForwardResponseTrailerHeader(ctx.Writer, md)
		ctx.Header(httpi.HeaderTransferEncoding, "chunked")
	}

	/*	st := HTTPStatusFromCode(se.Code)
		w.WriteHeader(st)*/
	if _, err := ctx.Writer.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
	if wantsTrailers {
		handleForwardResponseTrailer(ctx.Writer, md)
	}
}

type responseBody interface {
	XXX_ResponseBody() interface{}
}


func ForwardResponseMessage(ctx *gin.Context,md runtime.ServerMetadata ,message proto.Message) {
	if res, ok := message.(*response.HttpResponse); ok {
		for k, v := range res.Header {
			ctx.Header(k, v)
		}
		ctx.Status(int(res.StatusCode))
	}
	if md.HeaderMD == nil {
		md.HeaderMD = grpc.ServerTransportStreamFromContext(ctx.Request.Context()).(*runtime.ServerTransportStream).Header()
	}

	handleForwardResponseServerMetadata(ctx.Writer, md)
	handleForwardResponseTrailerHeader(ctx.Writer, md)

	contentType := jsonpb.JsonPb.ContentType(message)
	ctx.Header(httpi.HeaderContentType, contentType)

	var buf []byte
	var err error
	if rb, ok := message.(responseBody); ok {
		buf, err = jsonpb.JsonPb.Marshal(rb.XXX_ResponseBody())
	} else {
		buf, err = jsonpb.JsonPb.Marshal(message)
	}
	if err != nil {
		grpclog.Infof("Marshal error: %v", err)
		HTTPError(ctx, err)
		return
	}

	if _, err = ctx.Writer.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

	handleForwardResponseTrailer(ctx.Writer, md)
}
