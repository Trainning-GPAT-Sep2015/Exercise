package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type LogWriter struct {
	http.ResponseWriter
	code int
}

func (this *LogWriter) WriteHeader(code int) {
	this.code = code
	this.ResponseWriter.WriteHeader(code)
}

func LogMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			l := &LogWriter{ResponseWriter: w}
			h.ServeHTTP(l, r)
			responseTime := time.Now().Sub(start)
			log.Println(r.Method, r.RequestURI, "status", l.code, "duration", responseTime)
		})
}

func RecoverMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				e := recover()
				if e != nil {
					log.Println("RECOVER", e)
				}
			}()
			h.ServeHTTP(w, r)
		})

}

func AuthMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.RequestURI, "secret") {
				w.WriteHeader(403)
				fmt.Fprintf(w, "Not allow")
				return
			}
			h.ServeHTTP(w, r)
		})
}
