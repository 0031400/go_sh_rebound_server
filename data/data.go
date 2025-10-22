package data

import "github.com/gorilla/websocket"

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

func AddNode() NodeInfo {
	newNode := NodeInfo{Id: Id}
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
