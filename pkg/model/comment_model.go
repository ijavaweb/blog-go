package model
type ArticleComment struct {
	Id int32 `gorm:"column:id" ,json:"id"`
	CommentContent int32 `gorm:"column:comment_content" ,json:"comment_content"`
	CreateTime int32 `gorm:"column:create_time" ,json:"create_time"`
}
