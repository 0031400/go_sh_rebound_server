package main

import (
	"go_sh_rebound_server/config"
	"go_sh_rebound_server/logger"
	"go_sh_rebound_server/router"
	"log"
	"net/http"
)

func main() {
	logger.Init()
	config.Init()
	r := router.SetupRouter()
	log.Println("The server is listening on " + config.Addr)
	err := http.ListenAndServe(config.Addr, r)
	if err != nil {
		log.Panicln(err)
	}
}
