package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

var flArticleDir = flag.String("dir", "", "Aticle directory")

type Article struct {
	path string
	name string
}

var gArticles []Article

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)

	fmt.Fprintf(w, "<h1>Welcome to my blog</h1>")
	fmt.Fprintf(w, "<ul>")
	for _, article := range gArticles {
		fmt.Fprintf(w, `<li><a href="/article/%v">%v</a></li>`, article.name, article.name)
	}
	fmt.Fprintf(w, "</ul>")
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	uri := req.RequestURI
	name := uri[len("/article/"):]
	var selected Article

	for _, article := range gArticles {
		if name == article.name || name+".md" == article.name {
			selected = article
		}
	}

	if selected.path == "" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Not found")
		return
	}

	data, err := ioutil.ReadFile(selected.path)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	unsafe := blackfriday.MarkdownCommon(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write(html)
}

func loadAllArticles(root string) ([]Article, error) {
	articles := make([]Article, 4)[:0]
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			name := filepath.Base(path)
			if filepath.Ext(name) == ".md" {
				articles = append(articles, Article{
					path: path,
					name: name,
				})
			}
			return nil
		})

	return articles, err
}

func main() {
	flag.Parse()
	articles, err := loadAllArticles(*flArticleDir)
	if err != nil {
		log.Fatalln("Error reading articles", err)
	}

	gArticles = articles

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/article/", articleHandler)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
