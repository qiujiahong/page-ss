package screenShotService

import (
	"errors"
	"net/http"
	"page-ss/src/service/logger"
	"page-ss/src/utils"
)

func GetScreenShot(url string,quality int64,delay int64,forceUpdate bool) (error, []byte)  {

	logger.Log.Debug("proxy url = ",url)
	err,data := utils.GetFullScreenImageBytesWithHeader(url,quality,nil,nil,delay)
	if err != nil {
		return errors.New("image error"),nil
	}
	logger.Log.Debug("proxy url end return length  = ",len(data))

	return nil,data
}

func GetScreenShotWithHeader(url string,quality int64,headers map[string]interface{} ,cookies []*http.Cookie ,delay int64,forceUpdate bool) (error, []byte)  {
	logger.Log.Debug("proxy url = ",url)
	err,data := utils.GetFullScreenImageBytesWithHeader(url,quality,headers,cookies,delay)
	if err != nil {
		return errors.New("image error"),nil
	}
	logger.Log.Debug("proxy url end return length  = ",len(data))

	return nil,data
}