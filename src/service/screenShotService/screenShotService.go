package screenShotService

import (
	"errors"
	"net/http"
	"page-ss/src/service/logger"
	"page-ss/src/utils"
)

func GetScreenShot(url string,quality int64) (error, []byte)  {

	logger.Log.Debug("proxy url = ",url)
	err,data := utils.GetFullScreenImageBytes(url,quality)
	if err != nil {
		return errors.New("image error"),nil
	}
	logger.Log.Debug("proxy url end return length  = ",len(data))

	return nil,data
}

func GetScreenShotWithHeader(url string,quality int64,headers map[string]string ,cookies []*http.Cookie ) (error, []byte)  {
	logger.Log.Debug("proxy url = ",url)
	err,data := utils.GetFullScreenImageBytes(url,quality)
	if err != nil {
		return errors.New("image error"),nil
	}
	logger.Log.Debug("proxy url end return length  = ",len(data))

	return nil,data
}