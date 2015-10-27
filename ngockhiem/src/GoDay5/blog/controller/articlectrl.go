package controller

import (
	"GoDay5/blog/context"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Article struct {
	Path string
	Name string
}

var Articles []Article

func LoadArticles(root string) ([]Article, error) {
	articles := make([]Article, 4)[:0]
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		name := filepath.Base(path)
		if filepath.Ext(name) == ".md" {
			articles = append(articles, Article{
				Path: path,
				Name: name,
			})
		}
		return nil
	})
	return articles, err
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	name := context.GetContext(r).ByName("name")
	var article Article
	for _, a := range Articles {
		if a.Name == name || a.Name == name+".md" {
			article = a
			break
		}
	}
	if article.Path == "" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Article Not Found")
		return
	}

	content, err := ioutil.ReadFile(article.Path)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, " ERROR")
		return
	}

	unsafe := blackfriday.MarkdownCommon(content)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write(html)
}
