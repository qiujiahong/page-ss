package controllers

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"page-ss/src/config"
	"page-ss/src/utils"
	"strconv"
)

func (s  *Server) render(ctx *macaron.Context)   {
	fmt.Printf("path =  %v\r\n",ctx.Req.RequestURI)
	str := utils.Substr(ctx.Req.RequestURI,len("/render"),len(ctx.Req.RequestURI) -len("/render"))
	url := config.Global.ProxyUrl+str
	fmt.Printf("substr = %v \r\n",url)
	data := utils.GetFullScreenImageBytes(url,90,"data/newImage.png")
	fmt.Printf("step 2 = %d \r\n",len(data))
	ctx.Resp.WriteHeader(200)
	ctx.Resp.Header().Set("Content-Type", "image/png")
	ctx.Resp.Header().Set("C	ontent-Length", strconv.Itoa(len(data)))
	fmt.Printf("step 3 data length = %d \r\n", len(data))
	ctx.Resp.Write(data)
}
