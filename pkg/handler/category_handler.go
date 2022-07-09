package handler

import (
	"blog-go/pkg/logger"
	"blog-go/pkg/model"
	"blog-go/pkg/service"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context)  {
	category:=&model.Category{}
	err:=c.ShouldBindJSON(category)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	err=service.CreateCategory(category)
	if err!= nil {
		c.JSON(400,"failed")
		logger.ErrorLogger.Error(err)
		return
	}
	c.JSON(200,"success")
}

func UpdateCategory(c *gin.Context)  {
	category:=&model.Category{}
	err:=c.ShouldBindJSON(category)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	if category.Id<0 || category.Status!=1 {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"illegal param")
		return
	}
	err=service.UpdateCategory(category)
	if err!= nil {
		c.JSON(400,"failed")
		logger.ErrorLogger.Error(err)
		return
	}
	c.JSON(200,"success")
}

func DeleteCategory(c *gin.Context)  {
	category:=&model.Category{}
	err:=c.ShouldBindJSON(category)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	if category.Id<0 || category.Status!=1 {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"illegal param")
		return
	}
	err=service.DeleteCategory(category)
	if err!= nil {
		c.JSON(400,"failed")
		logger.ErrorLogger.Error(err)
		return
	}
	c.JSON(200,"success")
}

func GetCategoryList(c *gin.Context)  {
	data,err:=service.GetCategoryList()
	if err != nil {
		Base.responseError(c,err)
	}
	Base.responseOk(c,data)
}