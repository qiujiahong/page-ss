package main


import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
)


func main() {
	config.Init()
	fmt.Printf("config:%v",config.Global)
}
