package main

import (
	"blogrethink/querry"
	"fmt"
	r "github.com/dancannon/gorethink"
	"log"
)

var session *r.Session

func main() {

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "blog",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	authors, _ := querry.GetAllAuthor(session)
	articles, _ := querry.GetAllArticle(session)
	fmt.Println(authors)
	fmt.Println(articles)
	fmt.Println("DONE")
}
