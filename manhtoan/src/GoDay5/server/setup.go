package server

import (
	"fmt"
	"net/http"

	"GoDay5/api/rethinkdb"
	ctx "GoDay5/context"
	"GoDay5/dbscript"
	"GoDay5/handlers"
	"GoDay5/middlewares"

	"github.com/dancannon/gorethink"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type setupStruct struct {
	Config

	Rethink *rethinkdb.Instance
	Handler http.Handler
}

func setup(cfg Config) *setupStruct {
	s := &setupStruct{Config: cfg}
	s.setupRoutes()

	return s
}

func (s *setupStruct) setupRethink() {
	cfg := s.Config
	re, err := rethinkdb.NewInstance(gorethink.ConnectOpts{
		Address:  cfg.RethinkDB.Addr + ":" + cfg.RethinkDB.Port,
		Database: cfg.RethinkDB.DBName,
	})

	if err != nil {
		fmt.Println("Could not connect to RethinkDB")
	}

	s.Rethink = re

	script := dbscript.NewRethinkScript(s.Rethink, s.RethinkDB.DBName)
	if err := script.Setup(); err != nil {
		fmt.Println("Error generating data", err)
	}
}

func (s *setupStruct) setupRoutes() {

	auth := func(h http.HandlerFunc) httprouter.Handle {
		return ctx.HttpAdapter(ctx.Middleware(middlewares.AuthMiddleware(h)))
	}

	router := httprouter.New()

	{
		router.GET("/", handlers.Index)
	}

	{
		router.GET("/articles/:article", auth(handlers.ShowArticle))
	}

	s.Handler = context.ClearHandler(router)
}
