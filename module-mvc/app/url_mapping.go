package app

import (
	"module-mvc/controllers"
)

func mapUrls() {
	router.GET("/users", controllers.GetAllUserNew)
	router.GET("/users/:limit", controllers.GetAllUserNew)
	router.GET("/user/:user_id", controllers.GetUser)

	// html template output
	router.GET("/site/index", controllers.Index)

	// mysql
	router.GET("/site/popup-db", controllers.PopupDB)
}
