package db

import (
	"blog-go/pkg/model"
	"errors"
)

func CreateArticle(article *model.ArticleModel) (int32,error)  {
	db:=DB.Table("article").Create(article)
	return article.Id,db.Error
}

func UpdateArticle(article *model.ArticleModel)  error {
	db := DB.Table("article").Update(article).Where("id=?", article.Id)
	return db.Error
}

func DeleteArticle(article *model.ArticleModel)  error {
	db:=DB.Table("article").Update(article).Where("id=?",article.Id)
	return db.Error
}

func ListArticle(limit,offset int32) ([]model.ArticleModel,error)  {
	articles:=make([]model.ArticleModel,0)
	db:=DB.Table("article").Limit(limit).Offset(offset).Order("update_time desc").Find(&articles)
	if db.Error != nil {
		return nil,db.Error
	}
	return articles,nil
}
func GetArticleById(id int32) (*model.ArticleModel,error)  {
	if id < 0 {
		return nil, errors.New("invalid article id")
	}
	article:=&model.ArticleModel{}
	db:=DB.Table("article").Where("id=?",id).Find(article)
	if db.Error != nil {
		return nil,db.Error
	}
	return article,nil
}
func GetArticleCount() (*int32,error)  {
	var c int32
	db:=DB.Table("article").Count(&c)
	if db.Error != nil {
		return nil,db.Error
	}
	return &c,nil
}
func GetArticleCountByCategory(category string) (*int32,error)  {
	var c int32
	db:=DB.Table("article").Where("category=?",category).Count(&c)
	if db.Error != nil {
		return nil,db.Error
	}
	return &c,nil
}

func GetArticleByCategory(category string) ([]model.ArticleModel,error)  {
	articles:=make([]model.ArticleModel,0)
	db:=DB.Table("article").Where("category=?",category).Order("update_time desc").Find(&articles)
	if db.Error != nil {
		return nil,db.Error
	}
	return articles,nil
}