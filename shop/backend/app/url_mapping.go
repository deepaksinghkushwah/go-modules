package app

import (
	"shop/controllers"
	"shop/helpers"

	"github.com/gin-contrib/static"
)

func mapUrls() {
	router.Use(static.Serve("/", static.LocalFile("./static", false)))

	// api urls
	api := router.Group("/api")
	api.Use(helpers.Authorized)
	{
		api.GET("/users", controllers.GetAllUsers) // http://localhost:8082/api/users with params
		api.GET("/user", controllers.GetUser)
	}

	router.POST("/login", controllers.LoginHandler)       // http://localhost:8082/login with params
	router.POST("/register", controllers.RegisterHandler) // http://localhost:8082/register with params

}
