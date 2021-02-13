package main


import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
)


func main() {
	config.Init()
	fmt.Printf("config:%v",config.Global)

	m := macaron.Classic()
	m.Use(macaron.Logger())
	m.Get("/render/*", myHandler)
	m.Run()
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}