package data

import (
	"sync"

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

var mu sync.RWMutex
var NodeInfos []NodeInfo
var Id = 1

func AddNode(c *websocket.Conn, hostname string, addr string) NodeInfo {
	newNode := NodeInfo{Id: Id, C: c, Hostname: hostname, Addr: addr, ReadChan: make(chan []byte), WriteChan: make(chan []byte), ExitChan: make(chan struct{})}
	mu.Lock()
	NodeInfos = append(NodeInfos, newNode)
	mu.Unlock()
	Id++
	return newNode
}
func FindNode(id int) NodeInfo {
	mu.RLock()
	for _, v := range NodeInfos {
		if v.Id == id {
			mu.RUnlock()
			return v
		}
	}
	mu.RUnlock()
	return NodeInfo{Id: 0}
}
func DelNode(id int) {
	if id == 0 {
		return
	}
	for i, v := range NodeInfos {
		if v.Id == id {
			mu.Lock()
			NodeInfos = append(NodeInfos[:i], NodeInfos[i+1:]...)
			mu.Unlock()
			break
		}
	}
}
