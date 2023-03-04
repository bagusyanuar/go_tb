package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoute() *gin.Engine {
	route := gin.Default()

	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app_name": "teman_belajar",
			"version":  "1.0.0",
		})
	})
	return route
}
