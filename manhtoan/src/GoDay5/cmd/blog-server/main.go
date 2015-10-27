package main

import (
	"flag"
	"fmt"

	"GoDay5/server"
	"GoDay5/utils/load-config"
)

var flConfigFile = flag.String("config-file", "src/GoDay5/config-default.json", "Load config from file")

func main() {
	flag.Parse()

	var cfg server.Config
	err := loadConfig.FromFile(&cfg, *flConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	server.Start(cfg)
}
