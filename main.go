package main

import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
	"page-ss/src/utils"
)


func main() {
	config.Init()
	fmt.Printf("config:%v",config.Global)
	utils.GetFullScreenImage("https://www.baidu.com",90,"./data/fullScreenshot1.png")
	//controllers.Init()
}
