package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smiecj/gin_demo/controller"
	"github.com/smiecj/gin_demo/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()
	c := controller.NewJobController()
	{
		docs.SwaggerInfo.BasePath = ""
		docs.SwaggerInfo.Title = "swagger docs"
		docs.SwaggerInfo.Description = "swagger desc"
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/job/list", c.GetJobs)

	// swagger
	// refer: https://github.com/swaggo/swag#how-to-use-it-with-gin
	// url: /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// start
	r.Run("0.0.0.0:8000")
}
