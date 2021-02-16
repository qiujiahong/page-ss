package dto

import (
	"github.com/unknwon/com"
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
	// get config parameters
	urlParam.ParDelay = urlParam.ParamsInt64("__parDelay")
	urlParam.ForceUpdate =  urlParam.ParamsBool("__forceUpdate")
	urlParam.UseCache   = urlParam.ParamsBool("__useCache")
	urlParam.AutoFlush  = urlParam.ParamsBool("__autoFlush")
	urlParam.ValidityDays =urlParam.ParamsInt("__validityDays")
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