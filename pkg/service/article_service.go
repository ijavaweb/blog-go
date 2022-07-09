package service

import (
	"blog-go/pkg/db"
	"blog-go/pkg/model"
	"blog-go/pkg/util"
	"github.com/pkg/errors"
	"time"
)

func CreateArticle(article *model.Article) (*int32,error)  {
	if article==nil {
		return nil, errors.New("article illegal")
	}
	articleModel:=&model.ArticleModel{}
	articleModel.Title = article.Title
	articleModel.Content = article.Content
	articleModel.Category = article.Category
	articleModel.Status=1
	articleModel.CreateTime = time.Now().Unix()
	articleModel.UpdateTime = time.Now().Unix()
	id,err:=db.CreateArticle(articleModel)
	return &id,err
}

func UpdateArticle(article *model.Article) error  {
	if article==nil {
		return errors.New("article illegal")
	}
	articleModel:=&model.ArticleModel{}
	articleModel.Title = article.Title
	articleModel.Content = article.Content
	articleModel.Category = article.Category
	articleModel.UpdateTime = time.Now().Unix()
	err:=db.UpdateArticle(articleModel)
	return err
}

func DeleteArticle(article *model.ArticleModel) error  {
	if article.Id<0 {
		return errors.New("illegal article")
	}
	article.Status=0
	err:=db.DeleteArticle(article)
	return err
}

func ListArticle(req *model.ListArticleRequest) (*model.ListArticleResponse,error)  {
	resp:=&model.ListArticleResponse{}
	count,err:=db.GetArticleCount()
	if err != nil {
		return resp,err
	}
	articles,err:=db.ListArticle(req.Limit,req.Offset)
	if err != nil {
		return resp, err
	}
	articlesProcessed:=make([]model.ArticleResponse,0)
	for _,article:= range articles {
		a:=model.ArticleResponse{}
		a.Id = article.Id
		a.Title = article.Title
		a.Content = article.Content
		a.CreateTime = util.Time2Date(int64(article.CreateTime))
		a.UpdateTime = util.Time2Date(int64(article.UpdateTime))
		a.LetterNum = int32(len(article.Content))
		a.Category = article.Category
		articlesProcessed = append(articlesProcessed,a)
	}
	resp.Articles = articlesProcessed
	resp.Total = *count
	return resp,nil
}

func GetArticleById(id int32) (*model.ArticleResponse,error) {
	resp:=&model.ArticleResponse{}
	article,err:=db.GetArticleById(id)
	if err!= nil {
		return nil,err
	}
	resp.Id = article.Id
	resp.Title = article.Title
	resp.Content = article.Content
	resp.CreateTime = util.Time2Date(int64(article.CreateTime))
	resp.UpdateTime = util.Time2Date(int64(article.UpdateTime))
	resp.Category = article.Category
	resp.LetterNum = int32(len(article.Content))
	return resp,nil
}

func GetArticleByCategory(category string) (*model.ListArticleResponse,error) {
	resp:=&model.ListArticleResponse{}
	count,err:=db.GetArticleCountByCategory(category)
	if err != nil {
		return resp,err
	}
	articles,err:=db.GetArticleByCategory(category)
	if err != nil {
		return resp, err
	}
	articlesProcessed:=make([]model.ArticleResponse,0)
	for _,article:= range articles {
		a:=model.ArticleResponse{}
		a.Id = article.Id
		a.Title = article.Title
		a.Content = article.Content
		a.CreateTime = util.Time2Date(int64(article.CreateTime))
		a.UpdateTime = util.Time2Date(int64(article.UpdateTime))
		a.LetterNum = int32(len(article.Content))
		a.Category = article.Category
		articlesProcessed = append(articlesProcessed,a)
	}
	resp.Articles = articlesProcessed
	resp.Total = *count
	return resp,nil
}