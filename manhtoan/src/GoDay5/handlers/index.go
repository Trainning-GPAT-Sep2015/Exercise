package handlers

import (
	// "html/template"
	"GoDay5/utils/load-dir"
	"GoDay5/utils/load-markdown"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sourcePath := "src/GoDay5/database/index.md"

	html, err := loadMarkdown.FromFile(sourcePath)
	if err != nil {
		fmt.Println(err)
	}

	files, err := loadDir.ShowFileList("src/GoDay5/database")
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write(html)

	fmt.Fprintln(w, "<ul>")
	for _, value := range files {
		fmt.Fprintf(w, `<li><a href="http://localhost:3000/articles/%s">%s</a></li>`, value.Name(), value.Name())
	}
	fmt.Fprintln(w, "</ul>")
}
