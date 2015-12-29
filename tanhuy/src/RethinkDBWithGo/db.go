package main

import (
	"RethinkDBWithGo/Query"
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
		log.Fatal(err.Error())
	}
	// authors, _ := Query.GetAllAuthors(session)
	// //articles,_:=queryblog.GetAllArticles(session)
	// fmt.Println(authors)
	// //fmt.Println(articles)
	// fmt.Println("DONE")
	message, _ := Query.GetMessagesByRoom(session, 3)
	for _, item := range message {
		fmt.Println(item)
	}
	cursor, _ := Query.SubscribeMessages(session, 3, func(err error, item *Query.Message) {
		fmt.Printf("%v", item.Content)
	})
	ch := make(chan int)
	<-ch
	fmt.Println(cursor)
}
