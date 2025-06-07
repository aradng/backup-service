package api

import (
	"backup-service/api/backup"
	"backup-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
)

var App *gin.Engine
var Router *gin.RouterGroup

func init(){
	App = gin.Default()
	docs.SwaggerInfo.BasePath = "/api"
	Router = App.Group("/api")
	Router.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == docs.SwaggerInfo.BasePath + "/swagger/" {
			c.Redirect(302, docs.SwaggerInfo.BasePath + "/swagger/index.html")
			return
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})
	backup.RegisterRoutes(Router)
}