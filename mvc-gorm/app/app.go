package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	// when build binary for application, uncomment below line to enable production mode
	// gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("templates/*")
	router.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secrettop"))))
}

// StartApp function
func StartApp() {
	mapUrls()

	if err := router.Run(":8081"); err != nil {
		panic(err)
	}
}
