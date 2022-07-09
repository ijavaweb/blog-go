package model
type Category struct {
	Id int32 `gorm:"column:id" ,json:"id"`
	Name int32 `gorm:"column:name" ,json:"name"`
	CreateTime int32 `gorm:"column:create_time" ,json:"create_time"`
	Status int32 `gorm:"column:status" ,json:"status"`
}
type Result struct {
	Category string `gorm:"column:category" ,json:"category"`
}