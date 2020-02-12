package app

import (
	"shop/controllers"
	"shop/helpers"

	"github.com/gin-contrib/static"
)

func mapUrls() {
	router.Use(static.Serve("/", static.LocalFile("./static", false)))
	// api type urls
	api := router.Group("/api")
	api.Use(helpers.Authorized)
	{
		api.GET("/users", controllers.GetAllUserNew)
		api.GET("/user", controllers.GetUser)
		api.POST("/logout", controllers.Logout) // http://localhost:8081/site/logout
	}

	router.POST("/login", controllers.LoginHandler)       // http://localhost:8081/site/login?username=test3&password=123456
	router.POST("/register", controllers.RegisterHandler) //

}
