package middlewares

import (
	"GoDay5/context"
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.Get(req)
		articleName := ctx.Params.ByName("article")

		if strings.Contains(articleName, "secret") {
			w.WriteHeader(403)
			fmt.Fprint(w, "Sorry, this is security site")
			return
		}
		next.ServeHTTP(w, req)
	})
}

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		next.ServeHTTP(w, req)
	})
}
