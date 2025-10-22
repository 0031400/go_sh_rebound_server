package router

import (
	"go_sh_rebound_server/handler"
	"go_sh_rebound_server/middleware"
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/nodes", middleware.AuthClientMiddleWare(handler.NodesHandler))
	mux.HandleFunc("/client/ws", middleware.AuthClientMiddleWare(handler.ClientWsHandler))
	mux.HandleFunc("/node/ws", middleware.AuthNodeMiddleWare(handler.NodeWsHandler))
	return mux
}
