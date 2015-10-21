package main

import (
	"fmt"
	"log"
	"net/http"
	"now"
	"runtime/debug"
	"strings"
)

type LogWriter struct {
	http.ResponseWriter
	code int
}

func (l *LogWriter) WriteHeader(code int) {
	l.code = code
	l.ResponseWriter.WriteHeader(code)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {

			log.Printf("START", req.RequestURI)
			start := time.Now()

			logWriter := &LogWriter{ResponseWiter: w}
			h.ServeHTTP(logWritter, req)

			ctx := getContext(req)
			name := ctx.Params.ByName("name")

			duration := time.Now().Sub(start)
			log.Printf("END %v %v %v %v", req.RequestURI, logWriter.code, duration, name)
		})
}

func recoverMiddleware(h http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {

			defer func() {
				e := recover()
				if e != nil {
					log.Printf("RECOVER %v", e)
					debug.PrintStack()
					fmt.Fprintf(w, "Oops, something was wrong with us!")
				}
			}()

			h.ServeHTTP(w, req)
		})
}

func authMiddleware(h http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {

			ctx := getContext(req)
			name := ctx.Params.ByName("name")
			if strings.Contains("secret") {
				w.WriteHeader(403)
				fmt.Fprintf(w, "%v", "Not allowed")
				return
			}
			h.ServeHTTP(w, req)
		})
}
