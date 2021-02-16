package controllers

import (
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/models/dto"
	"page-ss/src/service/logger"
	"page-ss/src/service/screenShotService"
	"page-ss/src/utils"
)

func (s  *Server) render(ctx *macaron.Context)   {
	// 1.处理url
	str := utils.Substr(ctx.Req.RequestURI,len("/render"),len(ctx.Req.RequestURI) -len("/render"))
	// 2.获取url参数
	urlParam := dto.UrlParam{}
	urlParam.Init(str)
	// 3.整理proxy截图url
	url := config.Global.ProxyUrl+str
	// 4.获取图片
	err,data :=	screenShotService.GetScreenShotWithHeader(url,90,nil,nil,urlParam)

	// 5. 应答
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
	// 1. 整理请求url
	str := utils.Substr(ctx.Req.RequestURI,len("/renderWithHeader"),len(ctx.Req.RequestURI) -len("/renderWithHeader"))
	// 2. 处理url参数
	urlParam := dto.UrlParam{}
	urlParam.Init(str)
	// 3. 整理proxy截图url
	url := config.Global.ProxyUrl+str
	// 4. 处理cookie
	headers := GetHeaders(ctx)
	cookies := ctx.Req.Cookies()
	// 5. 获取图片
	err,data :=	screenShotService.GetScreenShotWithHeader(url,90,headers,cookies,urlParam)

	// 6. 应答
	if err!=nil {
		ctx.Resp.WriteHeader(500)
		ctx.Resp.Write([]byte("get screenshot failed."))
		SendResponse(ctx,500,data,"")
	} else{
		logger.Log.Debug("proxy url end return length  = ",len(data))
		SendResponse(ctx,200,data,"image/png")
	}
}
