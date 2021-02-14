package main

import (
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
	"page-ss/src/controllers"
	"page-ss/src/service/logger"
)

func main() {
	setup()
	logger.Log.Info("Server started:")
	logger.Log.Info("config: ",config.Global)
	controllers.Init()
}

func setup()  {
	config.Init()
	logger.Init()
}

