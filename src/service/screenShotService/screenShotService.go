package screenShotService

import (
	"errors"
	"net/http"
	"page-ss/src/models/dto"
	"page-ss/src/service/logger"
	"page-ss/src/utils"
)



func GetScreenShotWithHeader(url string,quality int64,headers map[string]interface{} ,cookies []*http.Cookie ,urlParam dto.UrlParam) (error, []byte)  {
	logger.Log.Debug("proxy url = ",url)
	err,data := utils.GetFullScreenImageBytesWithHeader(url,quality,headers,cookies,urlParam.ParDelay)
	if err != nil {
		return errors.New("image error"),nil
	}
	logger.Log.Debug("proxy url end return length  = ",len(data))

	return nil,data
}