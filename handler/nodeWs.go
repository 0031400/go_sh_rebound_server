package handler

import (
	"go_sh_rebound_server/config"
	"go_sh_rebound_server/data"
	"log"
	"net/http"
	"slices"

	"github.com/gorilla/websocket"
)

func NodeWsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panicln("node ws upgrade fail", err)
	}
	defer func() {
		c.Close()
		log.Println("disconnect with node")
	}()
	thisId := 0
	defer func() {
		data.DelNode(thisId)
	}()
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Panicln(err)
	}
	if mt != websocket.BinaryMessage || !slices.Equal(message, []byte{0}) {
		log.Panicln("fail to handshake")
	}
	err = c.WriteMessage(websocket.TextMessage, []byte(config.NodeAuth))
	if err != nil {
		log.Panicln(err)
	}
	mt, message, err = c.ReadMessage()
	if err != nil {
		log.Panicln(err)
	}
	if mt != websocket.TextMessage {
		log.Println("fail to handshake")
	}
	node := data.AddNode(c, string(message), c.RemoteAddr().String())
	thisId = node.Id
	c.WriteMessage(websocket.BinaryMessage, []byte{0})
	log.Println("connect with node " + node.Addr)
	go func() {
		for msg := range node.WriteChan {
			err = c.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			go func() {
				node.ExitChan <- struct{}{}
			}()
			log.Panicln(err)
		}
		if mt == websocket.BinaryMessage {
			node.ReadChan <- message
		}
	}
}
