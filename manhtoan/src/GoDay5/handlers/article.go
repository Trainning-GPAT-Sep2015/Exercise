package handlers

import (
	"GoDay5/api/rethinkdb"
	"GoDay5/context"
	"GoDay5/utils/load-markdown"
	"fmt"
	// r "github.com/dancannon/gorethink"
	"net/http"
)

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Get(r)
	databasePath := "src/GoDay5/database" + "/" + ctx.Params.ByName("article")

	// Test connect database
	session := rethinkdb.Connect()
	rethinkdb.CreateTable(session)

	html, err := loadMarkdown.FromFile(databasePath)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "Not found")
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write(html)
}
