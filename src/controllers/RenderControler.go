package controllers

import (
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/service/logger"
	"page-ss/src/service/screenShotService"
	"page-ss/src/utils"
)

func (s  *Server) render(ctx *macaron.Context)   {
	str := utils.Substr(ctx.Req.RequestURI,len("/render"),len(ctx.Req.RequestURI) -len("/render"))
	url := config.Global.ProxyUrl+str
    err,data :=	screenShotService.GetScreenShot(url,90)
    if err!=nil {
		ctx.Resp.WriteHeader(500)
		ctx.Resp.Write([]byte("get screenshot failed."))
		SendResponse(ctx,500,data,"")

	} else{
		logger.Log.Debug("proxy url end return length  = ",len(data))
		SendResponse(ctx,200,data,"image/png")
		//ctx.Resp.WriteHeader(200)
		//ctx.Resp.Header().Set("Content-Type", "image/png")
		//ctx.Resp.Header().Set("Content-Length", strconv.Itoa(len(data)))
		//logger.Log.Debug("return data")
		//ctx.Resp.Write(data)
	}
}
