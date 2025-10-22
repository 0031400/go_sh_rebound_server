package handler

import (
	"encoding/json"
	"go_sh_rebound_server/data"
	"log"
	"net/http"
)

func NodesHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(data.NodeInfos)
	if err != nil {
		log.Panicln(err)
	}
	w.Write(b)
}
