package handler

import (
	"blog-go/pkg/logger"
	"blog-go/pkg/model"
	"blog-go/pkg/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context)  {
	user:=&model.UserInfo{}
	err:=c.ShouldBindJSON(user)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	err=service.CreateUser(user)
	if err!= nil {
		c.JSON(400,"failed")
		logger.ErrorLogger.Error(err)
		return
	}
	c.JSON(200,"success")
}
func UpdateUser(c *gin.Context)  {
	user:=&model.UserInfo{}
	err:=c.ShouldBindJSON(user)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	err=service.UpdateUser(user)
	if err!= nil {
		c.JSON(400,"failed")
		logger.ErrorLogger.Error(err)
		return
	}
	c.JSON(200,"success")
}
func DeleteUser(c *gin.Context)  {
	user:=&model.UserInfo{}
	err:=c.ShouldBindJSON(user)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	err=service.DeleteUser(user)
	if err!= nil {
		c.JSON(400,"failed")
		logger.ErrorLogger.Error(err)
		return
	}
	c.JSON(200,"success")
}