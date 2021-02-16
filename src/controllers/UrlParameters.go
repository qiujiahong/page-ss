package controllers

import (
	"github.com/unknwon/com"
	"strconv"
	"strings"
)

type UrlParam struct {
	params	map[string] string
}

func (urlParam *UrlParam) Init(path string)  {
	urlParam.params = make(map[string] string)
	temp :=  strings.Split(path,"?")
	if len(temp) < 2 {
		return
	}
	paramsArr := strings.Split(temp[1],"&")

	for _, s := range paramsArr {
		item :=  strings.Split(s,"=")
		if len(item) == 2{
			urlParam.params[item[0]] = item[1]
		}
	}
}

func (urlParam *UrlParam) GetParam(name string) string  {
	return ""
}

func (urlParam *UrlParam) ParamsInt64(name string) int64 {
	return com.StrTo(urlParam.params[name]).MustInt64()
}

func (urlParam *UrlParam) ParamsBool(name string) bool {
	ret,err := strconv.ParseBool(urlParam.params[name])
	if err != nil{
		return false
	}else {
		return  ret
	}
}