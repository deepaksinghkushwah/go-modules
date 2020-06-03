package admincontrollers

import (
	"log"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var userkey, loggedInKey string

func init() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln("Error at loading env file")
	}
	userkey = os.Getenv("userkey")
	loggedInKey = os.Getenv("loggedInKey")
}

//Dashboard will display log of blogs in admin
func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	c.HTML(200, "admin_dashboard.html", gin.H{
		"Title":      "Dashbaord Page",
		"IsLoggedIn": session.Get(loggedInKey),
		"flashes":    flash,
		"roleID":     session.Get("roleID").(uint),
	})
}
