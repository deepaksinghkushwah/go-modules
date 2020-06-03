package admincontrollers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//BlogList will display log of blogs in admin
func BlogList(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	c.HTML(200, "admin_blog_list.html", gin.H{
		"Title":      "Blog Page",
		"Content":    "This is sample content",
		"IsLoggedIn": session.Get(loggedInKey),
		"flashes":    flash,
		"roleID":     session.Get("roleID").(uint),
	})
}
