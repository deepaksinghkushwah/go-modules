package app

import (
	"module-mvc/controllers"
	"module-mvc/helpers"

	"github.com/gin-contrib/static"
)

func mapUrls() {
	router.Use(static.Serve("/", static.LocalFile("./static", false)))
	// api type urls
	api := router.Group("/api")
	// no middleware so no use of api.Use function
	api.GET("/users", controllers.GetAllUserNew)
	api.GET("/users/:limit", controllers.GetAllUserNew)
	api.GET("/user/:user_id", controllers.GetUser)

	// html template output
	router.GET("/", controllers.Index)

	// api type urls
	router.GET("/site/login", controllers.LoginForm)     // http://localhost:8081/site/login?username=test3&password=123456
	router.POST("/site/login", controllers.LoginHandler) // http://localhost:8081/site/login?username=test3&password=123456

	router.GET("/site/logout", controllers.Logout)  // http://localhost:8081/site/logout
	router.POST("/site/logout", controllers.Logout) // http://localhost:8081/site/logout

	router.GET("/site/register", controllers.RegisterForm)     //
	router.POST("/site/register", controllers.RegisterHandler) //

	// mysql
	router.GET("/site/popup-db", controllers.PopupDB) // http://localhost:8081/site/popup-db

	// member group with authorization required
	// authorized.Use(MiddleWareName) is used to apply middleware function which will be checked on each route before route executed
	authorized := router.Group("/member") // /member is prefix which will added on each route who is member of this route group
	authorized.Use(helpers.Authorized)
	{
		authorized.GET("/me", controllers.Me) // http://localhost:8081/member/me
		authorized.GET("/update-profile", controllers.UpdateProfile)
		authorized.POST("/update-profile", controllers.UpdateProfileHandler)
	}

	/**
	 * using groups we can seprate various routes with prefix like...
	 * v1 := router.Group("/v1")
	 * v1.Get("/test1", controllers.v1Test1) // http://localhost:8081/v1/test1,  or you can use other custom package instaed of controller like api
	 */
}
