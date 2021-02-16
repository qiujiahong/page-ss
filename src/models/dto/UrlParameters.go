package dto

import (
	"github.com/unknwon/com"
	"page-ss/src/service/logger"
	"sort"
	"strconv"
	"strings"
	"time"
)

type UrlParam struct {
	params	map[string] string
	ParDelay int64
	ForceUpdate bool
	UseCache  bool
	AutoFlush bool
	ValidityDays int
	UrlWithoutParams string
}

func (urlParam *UrlParam) Init(path string)  {
	urlParam.params = make(map[string] string)
	hostAndParams :=  strings.Split(path,"?")
	if len(hostAndParams) < 2 {
		return
	}
	paramsArr := strings.Split(hostAndParams[1],"&")
	
	for _, s := range paramsArr {
		item :=  strings.Split(s,"=")
		if len(item) == 2{
			urlParam.params[item[0]] = item[1]
		} else if len(item) == 1 {
			urlParam.params[item[0]] = "__nil__"
		}
	}
	urlParam.ForceUpdate =  urlParam.ParamsBool("__forceUpdate")
	delete(urlParam.params, "__forceUpdate")
	urlParam.ParDelay = urlParam.ParamsInt64("__parDelay")
	delete(urlParam.params, "__parDelay")
	urlParam.UseCache   = urlParam.ParamsBool("__useCache")
	delete(urlParam.params, "__useCache")
	urlParam.AutoFlush  = urlParam.ParamsBool("__autoFlush")
	delete(urlParam.params, "__autoFlush")
	urlParam.ValidityDays =urlParam.ParamsInt("__validityDays")
	delete(urlParam.params, "__validityDays")
	logger.Log.Debug("params11:  ",urlParam)
	var b strings.Builder
	b.Grow(50)
	b.WriteString(hostAndParams[0])
	if len(urlParam.params)>0{
		var keys []string
		for s := range urlParam.params {
			keys = append(keys, s)
		}
		sort.Strings(keys)
		b.WriteString("?")
		for _, k := range keys {
			if urlParam.params[k] == "__nil__"{
				b.WriteString(k)
			} else{
				b.WriteString(k+"="+urlParam.params[k])
			}
			b.WriteString("&")
		}
	    temp :=	b.String()
		urlParam.UrlWithoutParams =temp[:len(temp)-1]
	} else {
		urlParam.UrlWithoutParams =b.String()
	}
	//a[:len(a)-1]
	logger.Log.Debug("UrlWithoutParams = ",urlParam.UrlWithoutParams)
}

func (urlParam *UrlParam) GetParam(name string) string  {
	return ""
}

func (urlParam *UrlParam) ParamsInt64(name string) int64 {
	return com.StrTo(urlParam.params[name]).MustInt64()
}
func (urlParam *UrlParam) ParamsInt(name string) int {
	return com.StrTo(urlParam.params[name]).MustInt()
}

func (urlParam *UrlParam) ParamsBool(name string) bool {
	ret,err := strconv.ParseBool(urlParam.params[name])
	if err != nil{
		return false
	}else {
		return  ret
	}
}

func (urlParam *UrlParam) GetValidityDay() time.Time {
	day := urlParam.ValidityDays
	if day == 0{
		day = 90
	}
	return time.Now().AddDate(0,0,day)
}