package web_server

import (
	"blog-go/pkg/config"
	"blog-go/pkg/handler"
	"blog-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)
func StartWebServer() {
	r:=gin.Default()
	r.Use(Cors())
	Register(r)
	address := config.LoadConfig()
	go func() {
		err := r.Run(address)
		if err != nil {
			logger.ErrorLogger.WithField("Http Start err: ", err).WithError(err)
			panic(err)
		}
	}()
}
func Register(engine *gin.Engine)  {
	article:=engine.Group("/article")

	user:=engine.Group("/user")

	comment:=engine.Group("/comment")

	visitor:=engine.Group("/visitor")

	category:=engine.Group("/category")

	article.POST("/create",handler.CreateArticle)
	article.GET("/list", handler.ListArticle)
	article.PUT("/update",handler.UpdateArticle)
	article.DELETE("/delete",handler.DeleteArticle)
	article.GET("/:id", handler.GetArticleById)
	article.GET("/category/:category", handler.GetArticleByCategory)



	user.POST("/create",handler.CreateUser)
	user.PUT("/update",handler.UpdateUser)
	user.DELETE("/delete",handler.DeleteUser)

	comment.GET("/details/:articleId",handler.GetCommentById)
	comment.GET("/amount",handler.GetCommentAmount)

	visitor.GET("/amount/:articleId",handler.GetVisitorById)
	visitor.GET("/amount",handler.GetVisitorAmount)

	category.POST("/create",handler.CreateCategory)
	category.PUT("/update",handler.UpdateCategory)
	category.DELETE("/delete",handler.DeleteCategory)
	category.GET("/list",handler.GetCategoryList)

}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}