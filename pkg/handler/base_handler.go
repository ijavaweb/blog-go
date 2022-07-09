package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseHandler struct {}
var Base *BaseHandler

func (h *BaseHandler)responseOk( c *gin.Context, data interface{})  {
	result:=make(map[string]interface{})
	result["code"]=0
	result["msg"] = "success"
	result["data"] = data
	c.JSON(http.StatusOK,result)
}

func (h *BaseHandler)responseError( c *gin.Context, err error)  {
	result:=make(map[string]interface{})
	result["code"]=1
	result["msg"] = "success"
	result["data"] = err.Error()
	c.JSON(http.StatusOK,result)
}

func (h *BaseHandler)responseBadRequest( c *gin.Context, err error)  {
	result:=make(map[string]interface{})
	result["code"]=2
	result["msg"] = "success"
	result["data"] = err.Error()
	c.JSON(http.StatusOK,result)
}