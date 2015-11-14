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
	// authors, _ := querry.GetAllAuthor(session)
	// articles, _ := querry.GetAllArticle(session)
	// articlesfromtag, _ := querry.GetAllArticleByTag("the", session)
	// fmt.Printf("%v", authors)
	// fmt.Println(articles)
	// fmt.Println(articlesfromtag)
	// fmt.Println("DONE")

	message, _ := querry.GetMessagesByRoom(session, 3)
	for _, item := range message {
		fmt.Printf("%v : %v \n", item.CreatedTime, item.Content)
	}
	cursor, _ := querry.SubscribeMessages(session, 3, func(err error, item *querry.Message) {
		fmt.Printf("%v : %v \n", item.CreatedTime, item.Content)
	})
	ch := make(chan int)
	<-ch
	fmt.Println(cursor)
}
