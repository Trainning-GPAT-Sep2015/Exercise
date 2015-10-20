package server

import (
	"fmt"
	"net/http"
)

func Server() {

	http.HandlerFunc("/", mainHandlerAction)
	http.HandlerFunc("/articles", articlesHandlerAction)

	// Listen on a port
	http.ListenAndServe(":8080", nil)
}

func mainHandlerAction(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Welcome to my Personal Blog")
}

func articleHandlerAction(rw http.ResponseWriter, req *http.Request) {

}

func loadArticles() {

}
