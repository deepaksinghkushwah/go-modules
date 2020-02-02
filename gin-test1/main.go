package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello Wrod",
	})
}

func PostHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Post home page",
	})
}

func QuueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func PathParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostHomePage)
	r.GET("/query", QuueryStrings)        //query?name=dee&age=37
	r.GET("/path/:name/:age", PathParams) //query?name=dee&age=37

	r.Run()
}
