package config

import (
	"flag"
	"os"
)

var Addr, NodeAuth, ClientAuth string

func Init() {
	addr := flag.String("l", "", "the server listen addr")
	nodeAuth := flag.String("na", "", "the node authorization")
	clientAuth := flag.String("ca", "", "the client authorization")
	flag.Parse()
	addrEnv := os.Getenv("addr")
	nodeAuthEnv := os.Getenv("nodeAuth")
	clientAuthEnv := os.Getenv("clientAuth")
	if addrEnv != "" {
		*addr = "127.0.0.1:3000"
	} else if *addr == "" {
		*addr = "127.0.0.1:3000"
	}
	Addr = *addr
	if nodeAuthEnv != "" {
		*nodeAuth = nodeAuthEnv
	}
	NodeAuth = *nodeAuth
	if clientAuthEnv != "" {
		*clientAuth = clientAuthEnv
	}
	ClientAuth = *clientAuth
}
