package framework

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type Context struct{
	request           *http.Request
	responseWriter    http.ResponseWriter
	ctx               context.Context
	handler           ControllerHandler

	writerMux         *sync.Mutex
	hasTimeOut        bool

}

func NewContext(r *http.Request, w http.ResponseWriter) *Context{
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}

func (ctx *Context)WriteMux() *sync.Mutex{
	return ctx.writerMux
}

func (ctx *Context)GetRequest() *http.Request{
	return ctx.request
}

func (ctx *Context)GetResponse() http.ResponseWriter{
	return ctx.responseWriter
}

func (ctx *Context)SetHasTimeOut(){
	ctx.hasTimeOut = true
}

func (ctx *Context)HasTimeOut() bool{
	return ctx.hasTimeOut
}


func (ctx *Context) BaseContext() context.Context{
	return ctx.request.Context()
}

func (ctx *Context)Deadline()(deadline time.Time, ok bool){
	return ctx.BaseContext().Deadline()
}

func (ctx *Context)Done() <-chan struct{}{
	return ctx.BaseContext().Done()
}

func (ctx *Context)Err() error{
	return ctx.BaseContext().Err()
}

func (ctx *Context)Value(key interface{}) interface{}{
	return ctx.BaseContext().Value(key)
}

func (ctx *Context)QueryInt(key string,def int){
	
}