package main

import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
	"page-ss/src/server"
)


func main() {
	config.Init()
	fmt.Printf("config:%v",config.Global)
	server.Init()
}
