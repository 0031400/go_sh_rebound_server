package data

import (
	"github.com/gorilla/websocket"
)

type NodeInfo struct {
	C         *websocket.Conn `json:"-"`
	Id        int             `json:"id"`
	Hostname  string          `json:"hostname"`
	Addr      string          `json:"addr"`
	ReadChan  chan ([]byte)   `json:"-"`
	WriteChan chan ([]byte)   `json:"-"`
	ExitChan  chan (struct{}) `json:"-"`
}

var NodeInfos []NodeInfo
var Id = 1

func AddNode(c *websocket.Conn, hostname string, addr string) NodeInfo {
	newNode := NodeInfo{Id: Id, C: c, Hostname: hostname, Addr: addr, ReadChan: make(chan []byte), WriteChan: make(chan []byte), ExitChan: make(chan struct{})}
	NodeInfos = append(NodeInfos, newNode)
	Id++
	return newNode
}
func FindNode(id int) NodeInfo {
	for _, v := range NodeInfos {
		if v.Id == id {
			return v
		}
	}
	return NodeInfo{Id: 0}
}
func DelNode(id int) {
	if id == 0 {
		return
	}
	for i, v := range NodeInfos {
		if v.Id == id {
			NodeInfos = append(NodeInfos[:i], NodeInfos[i+1:]...)
			break
		}
	}
}
