package crawler

import (
	"context"
	"github.com/actliboy/hoper/server/go/lib/utils/generics/conctrl"
)

type Prop struct {
}

type Request = conctrl.Task[string, Prop]
type TaskMeta = conctrl.TaskMeta[string]
type TaskFunc = conctrl.TaskFunc[string, Prop]

func NewRequest(key string, kind conctrl.Kind, taskFunc TaskFunc) *Request {
	return &Request{
		TaskMeta: TaskMeta{
			Kind: kind,
			Key:  key,
		},
		TaskFunc: taskFunc,
	}
}

type Engine = conctrl.Engine[string, Prop, Prop]

func NewEngine(workerCount uint) *conctrl.Engine[string, Prop, Prop] {
	return conctrl.NewEngine[string, Prop, Prop](workerCount)
}

type HandleFunc func(ctx context.Context, url string) ([]*Request, error)

func NewUrlRequest(url string, handleFunc HandleFunc) *Request {
	if handleFunc == nil {
		return nil
	}
	return &Request{TaskMeta: TaskMeta{Key: url}, TaskFunc: func(ctx context.Context) ([]*Request, error) {
		return handleFunc(ctx, url)
	}}
}

func NewUrlKindRequest(url string, kind conctrl.Kind, handleFunc HandleFunc) *Request {
	if handleFunc == nil {
		return nil
	}
	req := NewUrlRequest(url, handleFunc)
	req.SetKind(kind)
	return req
}
