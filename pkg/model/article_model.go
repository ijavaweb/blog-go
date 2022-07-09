package model
type Article struct { //nolint:maligned
	Title string `json:"title"`
	Content string `json:"content"`
	Category string `json:"category"`
}
type ArticleModel struct {
	Id int32 `gorm:"column:id" ,json:"id"`
	Title string `gorm:"column:title" ,json:"title"`
	Content string `gorm:"column:content" ,json:"content"`
	CreateTime int64 `gorm:"column:create_time" ,json:"create_time"`
	UpdateTime int64 `gorm:"column:update_time" ,json:"update_time"`
	Category string `gorm:"column:category" ,json:"category"`
	Status int32 `gorm:"column:status" ,json:"status"`
}
type ArticleResponse struct {
	Id int32 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
	LetterNum int32 `json:"letter_num"`
	Category string `json:"category"`
}
type ListArticleRequest struct {
	Limit int32 `json:"limit"`
	Offset int32 `json:"offset"`
}
type ListArticleResponse struct {
	Articles []ArticleResponse `json:"articles"`
	Total int32 `json:"total"`
}