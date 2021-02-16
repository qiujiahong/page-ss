//+build linux darwin

package main

import (
	"encoding/json"
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
	"page-ss/src/controllers"
	"page-ss/src/models"
	"page-ss/src/service/logger"
	"syscall"
)

func main() {
	setup()
	logger.Log.Info("Server started:")
	//logger.Log.Info("config: ",config.Global)
	//logger.Log.Debug(config.Global)
	//logger.Log.Debug(c.DbConfig)
	controllers.Init()
}

func setup()  {
	config.Init()
	logger.Init()
	syscall.Umask(0)
	//utils.Umask()
	// 打印调试
	setting, _ := json.MarshalIndent(config.Global, "", "\t")
	logger.Log.Debug("config information:\r\n",string(setting))

	//设置db连接池
	models.Setup()
}

