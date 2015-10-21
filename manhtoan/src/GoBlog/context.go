package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var contextMap map[*http.Request]interface{}

type Context struct {
	Params httprouter.Params
}

func getContext(req *http.Request) *Context {
	if ctx, ok := contextMap[req]; ok {
		return ctx
	}

	ctx := &Context{}
	contextMap[req] = ctx
	return ctx
}

func clearContext(req *http.Request) {
	delete(contextMap, req)
}

func contextMiddleware(h http.Handler) http.Handler {

	return http.HandleFunc(
		func(w http.ResponseWriter, req *http.Request) {
			getContext(req)
			h()
			clearContext(req)
		})
}

func httpAdapter(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request, httproute.Params) {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		ctx := getContext(req)
		ctx.Params = ps

	}
}
