package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default())
	gin.ForceConsoleColor()
}

// StartApp function
func StartApp() {
	mapUrls()
	if err := router.Run(":8082"); err != nil {
		panic(err)
	}
}
