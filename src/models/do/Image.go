package do

import "time"


// Image 缓存图片
type Image struct {
	Id   int32     `gorm:"primaryKey;column:id;index"`
	Url string    `gorm:"uniqueIndex;column:url;size:200;comment:图片截图的url"`
	Path string    `gorm:"column:path;size:100;comment:图片存放的路径"`
	AutoFlush bool 		`gorm:"comment:是否自动刷新";json:"auto_flush"`		//是否自动刷新
	ValidityDay time.Time `gorm:"comment:有效时间";json:"validity_day"`
	CreateTime   time.Time `gorm:"default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`
	UpdateTime   time.Time `gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName 指定table的名字
func (Image) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "images"
}

