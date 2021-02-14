package controllers

import (
	"gopkg.in/macaron.v1"
	"strconv"
)

func SendResponse(ctx *macaron.Context, status int,data [] byte,contentType string)  {
	ctx.Resp.WriteHeader(status)
	ctx.Resp.Header().Set("Content-Length", strconv.Itoa(len(data)))
	if contentType != "" {
		ctx.Resp.Header().Set("Content-Type", contentType)
	}

	ctx.Resp.Write(data)
}