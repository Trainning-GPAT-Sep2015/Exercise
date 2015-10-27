package server

import (
	"fmt"
	"net/http"
)

type Config struct {
	Server struct {
		Addr string `json: "addr"`
		Port string `json: "port"`
	} `json: "server"`
}

func Start(cfg Config) {
	s := setup(cfg)

	listenAddr := cfg.Server.Addr + ":" + cfg.Server.Port

	fmt.Println("Blog server is running on: ", listenAddr)
	http.ListenAndServe(listenAddr, s.Handler)
}
