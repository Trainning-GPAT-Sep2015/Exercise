package controller

import (
	"blog/context"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Article struct {
	Path    string
	Name    string
	Content []byte
}

var Articles []Article

var articleTemplate = `
	<html>

	<head>
		<title>{{.Name}}</title>
	</head>

	<body>
	<a href="/">Home</a>
	<b>Content</b>
	<p>{{.Content}}</p>
	</body>

	</html>
`

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
	w.Header().Add("Content-Type", "text/html")
	ctx := context.GetContext(r)
	token, err := GetToken(r)
	if err != nil {
		w.WriteHeader(401)
		fmt.Fprintf(w, "Not Authorized")
		return
	}
	if Token(ctx.Token) != token {
		w.WriteHeader(401)
		fmt.Fprintf(w, "Not Authorized")
		return
	}

	name := ctx.ByName("name")
	fmt.Println(w.Header().Get("Token"))
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
	article.Content = html
	t := template.Must(template.New("Article").Parse(articleTemplate))
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	t.Execute(w, article)
}
