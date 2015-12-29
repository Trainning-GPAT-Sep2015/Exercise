package main

import (
	"blog/context"
	"blog/controller"
	"blog/middleware"
	"encoding/json"
	"flag"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Config struct {
	Articles string
	Host     string
	Port     int
}

var configFlg = flag.String("config", "", "config json path")

var homepage = `
<html>
<head>
<title>Blog</title>
</head>
<body>
<a href="/login">login</a>
<a href="/login/reset">reset</a>
<p><b>Welcom to my blog</b></p>
<ul>
{{ range .}}
<li><a href="/article/{{.}}">{{.}}</a></li>
{{ end }}
</ul>
</body>
</html>
`

func loadConfig(path string) (Config, error) {
	config := Config{}
	data, err1 := ioutil.ReadFile(path)
	if err1 != nil {
		return Config{}, err1
	}
	err2 := json.Unmarshal(data, &config)
	if err2 != nil {
		return Config{}, err2
	}
	return config, nil

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	var articleList []string
	for _, article := range controller.Articles {
		articleList = append(articleList, article.Name)
	}
	t := template.Must(template.New("Homepage").Parse(homepage))
	t.Execute(w, articleList)
}

func main() {
	flag.Parse()

	config, config_err := loadConfig(*configFlg)
	if config_err != nil {
		log.Fatalln(config_err)
	}

	articles, load_err := controller.LoadArticles(config.Articles)
	if load_err != nil {
		log.Fatalln(load_err)
	}

	controller.Articles = articles

	httpAdapter := func(f http.HandlerFunc) httprouter.Handle {
		return context.Adapter(f)
	}

	router := httprouter.New()
	router.GET("/", httpAdapter(rootHandler))
	router.GET("/article/:name", httpAdapter(controller.ArticlesHandler))
	router.GET("/login", httpAdapter(controller.OauthGet))
	router.GET("/login/reset", httpAdapter(controller.OauthReset))

	finalrouter := middleware.RecoverMiddleWare(middleware.AuthMiddleWare(context.ContextMiddleWare(middleware.LogMiddleWare(router))))

	log.Println("Listen and Serve at", config.Host, ":", strconv.Itoa(config.Port))
	log.Fatalln(http.ListenAndServe(config.Host+":"+strconv.Itoa(config.Port), finalrouter))

}
