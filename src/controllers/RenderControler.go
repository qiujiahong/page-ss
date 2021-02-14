package controllers

import (
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/service/logger"
	"page-ss/src/utils"
	"strconv"
)

func (s  *Server) render(ctx *macaron.Context)   {
	str := utils.Substr(ctx.Req.RequestURI,len("/render"),len(ctx.Req.RequestURI) -len("/render"))
	url := config.Global.ProxyUrl+str
	logger.Log.Debug("proxy url = ",url)
	data := utils.GetFullScreenImageBytes(url,90,"data/newImage.png")
	logger.Log.Debug("proxy url end return length  = ",len(data))
	ctx.Resp.WriteHeader(200)
	ctx.Resp.Header().Set("Content-Type", "image/png")
	ctx.Resp.Header().Set("Content-Length", strconv.Itoa(len(data)))
	logger.Log.Debug("return data")
	ctx.Resp.Write(data)
}
C