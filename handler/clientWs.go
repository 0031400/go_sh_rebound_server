package handler

import (
	"go_sh_rebound_server/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

func ClientWsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln("ws upgrade fail", err)
	}
	defer func() {
		c.Close()
		log.Println("disconnect from client")
	}()
	idStr := r.URL.Query()["id"]
	if idStr == nil {
		log.Println("lack id params")
		return
	}
	id, err := strconv.Atoi(idStr[0])
	if err != nil {
		log.Panicln("parse id fail", err)
	}
	theNode := data.FindNode(id)
	if theNode.Id == 0 {
		log.Panicln("find the node fail")
		return
	}
	theNode.WriteChan <- []byte{0}
	go func() {
		for d := range theNode.ReadChan {
			c.WriteMessage(websocket.BinaryMessage, d)
		}
	}()
	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			if mt == websocket.BinaryMessage {
				theNode.WriteChan <- message
			}
		}
	}()
	<-theNode.ExitChan
	log.Panicln("link with client stop because link of node stop")
}
