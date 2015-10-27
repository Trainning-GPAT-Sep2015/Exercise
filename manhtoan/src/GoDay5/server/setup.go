package server

import (
	"net/http"

	"github.com/gorilla/context"

	"GoDay5/handlers"
	"github.com/julienschmidt/httprouter"
)

type setupStruct struct {
	Config

	Handler http.Handler
}

func setup(cfg Config) *setupStruct {
	s := &setupStruct{Config: cfg}
	s.setupRoutes()

	return s
}

func (s *setupStruct) setupRoutes() {

	router := httprouter.New()

	{
		router.GET("/", handlers.Index)
	}

	{
		router.GET("/list", handlers.List)
	}

	s.Handler = context.ClearHandler(router)
}
