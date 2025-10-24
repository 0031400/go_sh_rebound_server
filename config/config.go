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
	if *addr == "" && addrEnv == "" {
		Addr = "127.0.0.1:3000"
	} else if addrEnv != "" {
		Addr = addrEnv
	} else {
		Addr = *addr
	}
	if *nodeAuth == "" && nodeAuthEnv != "" {
		NodeAuth = nodeAuthEnv
	} else {
		NodeAuth = *nodeAuth
	}
	if *clientAuth == "" && clientAuthEnv != "" {
		ClientAuth = clientAuthEnv
	} else {
		ClientAuth = *clientAuth
	}
}
