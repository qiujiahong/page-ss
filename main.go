package main


import (
	"fmt"
	_ "github.com/icattlecoder/godaemon"
	"page-ss/src/config"
	"gopkg.in/macaron.v1"
)


func main() {
	config.Init()
	fmt.Printf("config:%v",config.Global)

	m := macaron.Classic()
	m.Use(macaron.Logger())
	m.Get("/render/*", myHandler)
	//m.Get("/", func() string {
	//	return "Hello world!"
	//})
	m.Run()
}

func myHandler(ctx *macaron.Context) string {
	return "the request path is: " + ctx.Req.RequestURI
}