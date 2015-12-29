package context

import (
	// "fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var contextMap = make(map[*http.Request]*Context)

type Context struct {
	Params httprouter.Params
}

func Get(req *http.Request) *Context {
	if ctx, ok := contextMap[req]; ok {
		return ctx
	}
	ctx := &Context{}
	contextMap[req] = ctx
	return ctx
}

func Clear(req *http.Request) {
	delete(contextMap, req)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			Get(req)
			next.ServeHTTP(w, req)
			Clear(req)
		})
}

func HttpAdapter(next http.Handler) httprouter.Handle {
	return httprouter.Handle(
		func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
			ctx := Get(req)
			ctx.Params = ps
			next.ServeHTTP(w, req)
		})
}
