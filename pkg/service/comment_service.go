package service

import (
	"blog-go/pkg/db"
	"blog-go/pkg/model"
)

func GetCommentById(id int32) ([]model.ArticleComment,error)  {
	return db.GetCommentById(id)
}
func GetCommentAmount() (int32,error)  {
	return db.GetCommentAmount()
}