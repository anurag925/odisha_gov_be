package main

import (
	"net/http"
	"odisha_gov_be/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(server *gin.Engine) {
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})
	api := server.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			districtCtrl := controllers.DistrictCtrl{}
			v1.GET("/district", districtCtrl.Get)
			v1.POST("/district", districtCtrl.Post)
		}
	}
}
