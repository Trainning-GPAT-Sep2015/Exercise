package main

import (
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

func HandlerAdapter(f func(w http.ResponseWriter, r *http.Request)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		f(w, r)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("file")
	log.Println(handler.Filename)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := os.OpenFile("./image/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0755)
	defer file.Close()
	defer f.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte("http://localhost:4000/image/" + handler.Filename))
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Upload service"))
}
func main() {

	router := httprouter.New()
	router.POST("/upload", HandlerAdapter(UploadHandler))
	router.GET("/", HandlerAdapter(RootHandler))
	http.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.Handle("/", router)
	log.Println("Listen and Server at localhost:4000")
	log.Fatalln(http.ListenAndServe("localhost:4000", nil))
}
