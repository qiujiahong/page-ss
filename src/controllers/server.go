package controllers

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/service/logger"
	"time"
)

type Server struct {
	port int  //通讯的端口
}

var server = &Server{}

func Init () {
	server.port = config.Global.Port
	m := macaron.Classic()
	m.Use(macaron.Logger())
	m.Use(macaron.Static("images", macaron.StaticOptions{
		// Prefix is the optional prefix used to serve the static directory content. Default is empty string.
		Prefix: config.Global.Prefix+"/images",
		// SkipLogging will disable [Static] log messages when a static file is served. Default is false.
		SkipLogging: true,
		// IndexFile defines which file to serve as index if it exists. Default is "index.html".
		IndexFile: "index.html",
		// Expires defines which user-defined function to use for producing a HTTP Expires Header. Default is nil.
		// https://developers.google.com/speed/docs/insights/LeverageBrowserCaching
		Expires: func() string {
			return time.Now().Add(24 * 60 * time.Minute).UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
		},
	}))
	m.Get(config.Global.Prefix+"/", server.home)
	m.Get(config.Global.Prefix+"/render", server.render)
	m.Get(config.Global.Prefix+"/render/*", server.render)
	m.Get(config.Global.Prefix+"/renderWithHeader", server.renderWithHeader)
	m.Get(config.Global.Prefix+"/renderWithHeader/*", server.renderWithHeader)
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

