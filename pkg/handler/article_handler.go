package handler

import (
	"blog-go/pkg/logger"
	"blog-go/pkg/model"
	"blog-go/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func CreateArticle(c *gin.Context)  {
	article:=&model.Article{}
	err:=c.BindJSON(article)
	fmt.Println(*article)
	if err!=nil {
		logger.ErrorLogger.WithField("err",err).WithError(err)
		Base.responseError(c,err)
		return
	}
	id,err:=service.CreateArticle(article)
	if err!= nil {
		Base.responseError(c,err)
		logger.ErrorLogger.Error(err)
		return
	}
	data :=make(map[string]interface{})
	data["article_id"] = *id
	Base.responseOk(c,data)
}

func UpdateArticle(c *gin.Context)  {
	article:=&model.Article{}
	err:=c.ShouldBindJSON(article)
	if err!=nil {
		Base.responseError(c,err)
		c.JSON(400,"Bad Request")
		return
	}
	err=service.UpdateArticle(article)
	if err!= nil {
		Base.responseError(c,err)
		logger.ErrorLogger.Error(err)
		return
	}
	data :=make(map[string]interface{})
	data["article_id"] = 1
	Base.responseOk(c,data)
}

func DeleteArticle(c *gin.Context)  {
	article:=&model.ArticleModel{}
	err:=c.ShouldBindJSON(article)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		Base.responseError(c,err)
		return
	}
	err=service.DeleteArticle(article)
	if err!= nil {
		Base.responseError(c,err)
		logger.ErrorLogger.Error(err)
		return
	}
	data :=make(map[string]interface{})
	data["article_id"] = 1
	Base.responseOk(c,data)
}

func ListArticle(c *gin.Context) {
	req:=&model.ListArticleRequest{}
	limit,err:= strconv.Atoi(c.Query("limit"))
	offset,err:= strconv.Atoi(c.Query("offset"))
	req.Limit = int32(limit)
	req.Offset = int32(offset)
	if err != nil {
		Base.responseError(c,err)
		return
	}
	resp,err:=service.ListArticle(req)
	if err != nil {
		Base.responseError(c,err)
		return
	}
	Base.responseOk(c,resp)
}

func GetArticleById(c *gin.Context)  {
	id,err:= strconv.Atoi(c.Param("id"))
	if err != nil {
		Base.responseError(c,err)
		return
	}
	result,err := service.GetArticleById(int32(id))
	if err != nil {
		Base.responseError(c,err)
		return
	}
	Base.responseOk(c,result)
}

func GetArticleByCategory(c *gin.Context) {
	category := c.Param("category")
	resp,err:=service.GetArticleByCategory(category)
	if err != nil {
		Base.responseError(c,err)
		return
	}
	Base.responseOk(c,resp)
}