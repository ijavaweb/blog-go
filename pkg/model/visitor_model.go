package model
type Visitor struct {
	Id int32 `gorm:"column:id" ,json:"id"`
	Visitor int32 `gorm:"column:visitor" ,json:"visitor"`
}

