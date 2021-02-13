package controllers

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
)

type Server struct {
	port int  //通讯的端口
}

var server = &Server{}

func Init () {
	server.port = config.Global.Port
	m := macaron.Classic()
	m.Use(macaron.Logger())
	m.Get("/", server.home)
	m.Get("/render", server.render)
	m.Get("/render/*", server.render)
	m.Run(server.port)
}


func (s  *Server) home(ctx *macaron.Context)   string {
	return fmt.Sprintf("the request path is: %v,port is:%v",ctx.Req.RequestURI,  s.port)
}

func (s  *Server) render(ctx *macaron.Context)   string {
	return fmt.Sprintf("the request path is: %v,port is:%v",ctx.Req.RequestURI,  s.port)
}




