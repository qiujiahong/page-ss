package controllers

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/service/logger"
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
	m.Get("/renderWithHeader", server.renderWithHeader)
	m.Get("/renderWithHeader/*", server.renderWithHeader)
	m.Run(server.port)
}


func (s  *Server) home(ctx *macaron.Context)   string {
	GetCookies(ctx)
	return fmt.Sprintf("the request path is: %v,port is:%v",ctx.Req.RequestURI,  s.port)
}

// curl -H "zhonngguo:nick" -H "origin: https://stackoverflow.com"  http://localhost:8080
func GetHeaders(ctx *macaron.Context) map[string]interface{}  {
	var headers map[string]interface{}  = make(map[string]interface{})
	for s2, strings := range ctx.Req.Header {
		//logger.Log.Info("header:",s2,":",strings)
		headers[s2] = strings[0]
	}
	ctx.Req.Cookies()
	return  headers
}

//  curl http://localhost:8080 --cookie "user=root;pass=123456"

func GetCookies(ctx *macaron.Context) map[string]string  {
	//var headers map[string]string  = make(map[string]string)
	for i, cookie := range ctx.Req.Cookies() {
		logger.Log.Info(i,cookie,cookie.Domain)
	}
	return  nil
}

