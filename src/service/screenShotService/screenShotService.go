package screenShotService

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"page-ss/src/models"
	"page-ss/src/models/do"
	"page-ss/src/models/dto"
	"page-ss/src/service/logger"
	"page-ss/src/utils"
)


func GetScreenShotWithHeader(url string, quality int64, headers map[string]interface{} ,
	cookies []*http.Cookie , urlParam dto.UrlParam) (error, []byte)  {

	logger.Log.Debug("proxy url = ",url)
	var data []byte
	var err error = nil
	getNewImage := false
	image := do.Image{}
	db := models.GetDB()
	db.Where(&do.Image{Url: url}).First(&image)

	if urlParam.ForceUpdate { //强制更新
		err,data = utils.GetFullScreenImageBytesWithHeader(url,quality,headers,cookies,urlParam.ParDelay)
		if err != nil {
			return errors.New("image error"),nil
		}
		getNewImage = true
	} else {  // 检查缓存
		if image.Id != 0 {
			logger.Log.Debugf("image exist: %v %v",image.GetFullPath(),image.Url)
			data,err  = ioutil.ReadFile(image.GetFullPath())
			if err != nil {
				logger.Log.Error("get image cache error")
				err,data = utils.GetFullScreenImageBytesWithHeader(url,quality,headers,cookies,urlParam.ParDelay)
				if err != nil {
					return errors.New("image error"),nil
				}
				getNewImage = true
			}
		} else{
			logger.Log.Debug("no cache:",url)
			err,data = utils.GetFullScreenImageBytesWithHeader(url,quality,headers,cookies,urlParam.ParDelay)
			if err != nil {
				return errors.New("image error"),nil
			}
			getNewImage = true
		}
	}

	if getNewImage && urlParam.UseCache {
		logger.Log.Debug("update cache.",image)
		path,name := utils.GetImagePath()
		image.Url = url
		image.Path =path
		image.Name =name
		image.AutoFlush = urlParam.AutoFlush
		image.ValidityDay = urlParam.GetValidityDay()
		if image.Id == 0 { //插入
			result := db.Create(&image)
			if result.Error != nil {
				logger.Log.Error("insert record failed:",result.Error)
			}
		} else { // 更新
			db.Save(&image)
		}
		logger.Log.Debug("save files:",image.Path + image.Name)
		err := os.MkdirAll(image.Path, 0766)
		if err != nil {
			logger.Log.Error("create path failed",err)
		}
		if err := ioutil.WriteFile(image.Path+image.Name, data, 0640); err != nil {
			log.Fatal("save files error",err)
		}
	}

	logger.Log.Debug("proxy url end return length  = ",len(data))
	return err,data
}