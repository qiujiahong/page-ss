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
	}
}

func (s  *Server) renderWithHeader(ctx *macaron.Context)   {
	str := utils.Substr(ctx.Req.RequestURI,len("/renderWithHeader"),len(ctx.Req.RequestURI) -len("/renderWithHeader"))
	url := config.Global.ProxyUrl+str

	headers := GetHeaders(ctx)
	cookies := ctx.Req.Cookies()

	err,data :=	screenShotService.GetScreenShotWithHeader(url,90,headers,cookies)
	if err!=nil {
		ctx.Resp.WriteHeader(500)
		ctx.Resp.Write([]byte("get screenshot failed."))
		SendResponse(ctx,500,data,"")

	} else{
		logger.Log.Debug("proxy url end return length  = ",len(data))
		SendResponse(ctx,200,data,"image/png")
	}
}
