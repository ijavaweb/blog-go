package model

type UserInfo struct {
	Id int32 `gorm:"column:id" ,json:"id"`
	UserName string `gorm:"column:username" ,json:"username"`
	Password string `gorm:"column:password" ,json:"password"`
	Role string `gorm:"column:role" ,json:"role"`
	Status int32 `gorm:"column:status" ,json:"status"`
}