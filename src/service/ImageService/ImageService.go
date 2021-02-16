package ImageService

import (
	"page-ss/src/models"
	"page-ss/src/models/do"
)


// 检查 是否有缓存图片
// 如果有缓存则读文件返回字节流
// 如果没有缓存则重新截图
func RequestImageService(url string, data []byte ) ([]byte,error)  {
	db := models.GetDB()
	image := do.Image{}

	if err := db.Where(&do.Image{Url: url}).First(&image).Error; err != nil {
		// 处理错误...
	}


	// 1.查询数据库里面是否存在该url的图片
	// 2.如果存在则返回该条记录
	// 3.如果不存在则返回错误
	return  nil,nil
}