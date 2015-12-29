package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("starting server on http://localhost:8888/\nvalue is %s", value)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe(":8888", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file1")
	if err != nil {
		log.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(handler.Filename, data, 0777)
	if err != nil {
		log.Println(err)
	}
}
