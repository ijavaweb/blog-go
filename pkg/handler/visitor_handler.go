package handler

import (
	"blog-go/pkg/logger"
	"blog-go/pkg/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetVisitorById(c *gin.Context)  {
	articleId:=c.Param("articleId")
	id,err:=strconv.Atoi(articleId)
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"Bad Request")
		return
	}
	result,err:=service.GetVisitorById(int32(id))
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"failed")
		return
	}
	c.JSON(200,result)
}
func GetVisitorAmount(c *gin.Context)  {
	result,err:=service.GetVisitorAmount()
	if err!=nil {
		logger.ErrorLogger.Error(err)
		c.JSON(400,"failed")
		return
	}
	c.JSON(200,result)
}