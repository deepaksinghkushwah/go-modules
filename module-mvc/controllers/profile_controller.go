package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Me function display after logged in
func Me(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	currentPath := c.Request.URL.Path

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Title":       "My Status",
		"Content":     nil,
		"IsLoggedIn":  session.Get(loggedInKey),
		"flashes":     flash,
		"currentPath": currentPath,
	})

}
