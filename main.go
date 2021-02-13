package main

import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
	"page-ss/src/controllers"
)


func main() {
	fmt.Printf("abc")
	config.Init()
	fmt.Printf("config:%v",config.Global)
	//utils.GetFullScreenImage("https://www.baidu.com",90,"./data/fullScreenshot1.png")
	controllers.Init()
}
