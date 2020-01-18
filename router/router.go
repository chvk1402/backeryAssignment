package router

import (
	"backery/handlers"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	r.POST("/price", handlers.Calculate)
	return r
}
