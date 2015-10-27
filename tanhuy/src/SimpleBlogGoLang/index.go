package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Article struct {
	path string
	name string
}

type Config struct {
	Path string `json: "path"`
}

var cfg, _ = config()

func config() (Config, error) {
	var cfg = Config{}
	configFile, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Config file not found")
	}
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&cfg); err != nil {
		log.Fatal("Cannot parse config file")
	}
	defer configFile.Close()
	return cfg, err
}

func readArticle(filename string) ([]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %v %v", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf("Completed %v %v in %v", r.Method, r.URL.Path, time.Since(start))
	}
}
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		filename := vars["name"]
		if filename == "secret.md" {
			w.WriteHeader(403)
			fmt.Fprintf(w, "You are not allow to access this page")
		} else {
			next.ServeHTTP(w, r)
		}
	}
}

func listAllAritcles(root string) ([]Article, error) {
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

func articleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["name"]
	if filename == "" {
		w.WriteHeader(404)
		fmt.Fprintf(w, "404 Not found")
	}
	articles, _ := listAllAritcles(cfg.Path)
	article := Article{}
	for _, val := range articles {
		if filename == val.name {
			article = val
		}
	}
	data, err := readArticle(article.path)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "Internal Error 500")
	}
	unsafe := blackfriday.MarkdownCommon(data)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write(html)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(200)
	fmt.Fprintf(w, "<h1>Hello, welcome to my Blog. My name is Julian</h1>")
	fmt.Fprintf(w, "<ul>")
	lsArticle, _ := listAllAritcles(cfg.Path)
	for _, article := range lsArticle {
		fmt.Fprintf(w, `<li><a href="article/%s">%s</a></li>`, article.name, article.name)
		// data, _ := readArticle(article.path)
	}
	fmt.Fprintf(w, "</ul>")

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", loggerMiddleware(rootHandler))
	router.HandleFunc("/article/{name}", authMiddleware(loggerMiddleware(articleHandler)))
	log.Fatalln(http.ListenAndServe(":8080", router))
}
