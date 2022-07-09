package db

import "blog-go/pkg/model"

func GetCommentById(id int32) ([]model.ArticleComment,error){
	var comments []model.ArticleComment
	db:=DB.Table("article_comment").Find(comments).Where("id=?",id)
	return comments,db.Error
}

func GetCommentAmount() (int32,error){
	count:=0
	db:=DB.Table("article_comment").Count(&count)
	return int32(count),db.Error
}
