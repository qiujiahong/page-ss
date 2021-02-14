package controllers

import (
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/service/logger"
	"page-ss/src/service/screenShotService"
	"page-ss/src/utils"
	"strconv"
)

func (s  *Server) render(ctx *macaron.Context)   {
	str := utils.Substr(ctx.Req.RequestURI,len("/render"),len(ctx.Req.RequestURI) -len("/render"))
	url := config.Global.ProxyUrl+str
    err,data :=	screenShotService.GetScreenShot(url,90)
    if err!=nil {
		ctx.Resp.WriteHeader(500)
		ctx.Resp.Write([]byte("get screenshot failed."))
	} else{
		logger.Log.Debug("proxy url end return length  = ",len(data))
		ctx.Resp.WriteHeader(200)
		ctx.Resp.Header().Set("Content-Type", "image/png")
		ctx.Resp.Header().Set("Content-Length", strconv.Itoa(len(data)))
		logger.Log.Debug("return data")
		ctx.Resp.Write(data)
	}
}
