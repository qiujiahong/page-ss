package do

import "time"


// Image 缓存图片
type Image struct {
	Id   int32     `gorm:"primaryKey;column:id;index"`
	Url string    `gorm:"index;column:url;size:200;comment:图片截图的url"`
	Path string    `gorm:"column:path;size:100;comment:图片存放的路径"`
	AutoFlush bool 		`gorm:"comment:是否自动刷新";json:"auto_flush"`		//是否自动刷新
	ValidityDay time.Time `gorm:"comment:有效时间";json:"validity_day"`
	CreateTime   time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime   time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
}

//https://gorm.io/docs/models.html