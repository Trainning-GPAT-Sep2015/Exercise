package context

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var ContextMap map[*http.Request]*Context = make(map[*http.Request]*Context)

type Context struct {
	httprouter.Params
	Token string
}

func GetContext(r *http.Request) *Context {
	if ctx, ok := ContextMap[r]; ok {
		return ctx
	}
	ctx := &Context{}
	ContextMap[r] = ctx
	return ctx
}

func ClearContext(r *http.Request) {
	delete(ContextMap, r)
}

func ContextMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := GetContext(r)
			ctx.Token = r.Header.Get("Authentication")
			fmt.Println(r.Header)
			h.ServeHTTP(w, r)
			ClearContext(r)
		})
}

func Adapter(f func(http.ResponseWriter, *http.Request)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := GetContext(r)
		ctx.Params = ps
		f(w, r)
	}
}
