package handler

import (
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
		log.Panicln(err)
	}
	defer c.Close()
	thisId := data.Id
	data.Id++
	defer func() {
		for i, v := range data.NodeInfos {
			if v.Id == thisId {
				data.NodeInfos = append(data.NodeInfos[:i], data.NodeInfos[i+1:]...)
				break
			}
		}
	}()
	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Panicln(err)
	}
	if mt != websocket.BinaryMessage || !slices.Equal(message, []byte{0}) {
		log.Panicln("fail to handshake")
		return
	}
	mt, message, err = c.ReadMessage()
	if err != nil {
		log.Panicln(err)
	}
	if mt != websocket.TextMessage {
		log.Println("fail to handshake")
		return
	}
	var node data.NodeInfo
	node.Hostname = string(message)
	node.Addr = c.RemoteAddr().String()
	node.Id = thisId
	node.C = c
	node.ReadChan = make(chan []byte)
	node.WriteChan = make(chan []byte)
	node.ExitChan = make(chan struct{})
	data.NodeInfos = append(data.NodeInfos, node)
	c.WriteMessage(websocket.BinaryMessage, []byte{0})
	log.Println("link from node " + node.Addr)
	go func() {
		for msg := range node.WriteChan {
			c.WriteMessage(websocket.BinaryMessage, msg)
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
