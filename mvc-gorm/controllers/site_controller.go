package controllers

import (
	"log"

	"mvc-gorm/helpers/general"
	"mvc-gorm/models"
	"net/http"
	"os"
	"strings"

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

// Index function
func Index(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	var roleID uint
	roleID = session.Get("roleID").(uint)

	c.HTML(200, "index.html", gin.H{
		"Title":      "Index Page",
		"Content":    "This is sample content",
		"IsLoggedIn": session.Get(loggedInKey),
		"flashes":    flash,
		"roleID":     roleID,
	})

}

// LoginForm display login form
func LoginForm(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	c.HTML(http.StatusOK, "login.html", gin.H{
		"Title":      "Login",
		"Content":    nil,
		"IsLoggedIn": session.Get(loggedInKey),
		"flashes":    flash,
		"roleID":     session.Get("roleID"),
	})

}

//LoginHandler to handle login request
func LoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim("username", " ") == "" || strings.Trim("password", " ") == "" {
		session.AddFlash("Username or password can not be empty")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/site/login")
	} else {
		db := general.GetDB()
		defer general.CloseDB(db)
		user := models.User{}
		db.Where("username = ?", username).First(&user)
		//fmt.Println("--------------------------")
		//log.Println(user.RoleID)
		//fmt.Println("--------------------------")
		if user.ID <= 0 {
			session.AddFlash("User not found")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/site/login")
		} else {
			if general.CheckPasswordHash(password, user.Password) {
				session.Set(userkey, username)
				session.Set("userID", user.ID)
				session.Set(loggedInKey, true)
				session.Set("roleID", uint(user.RoleID))
				if err := session.Save(); err != nil {
					session.AddFlash("Unauthorized Access")
					session.Save()
					c.Redirect(http.StatusSeeOther, "/site/login")
				}
				session.AddFlash("Logged in successfully")
				session.Save()
				c.Redirect(http.StatusSeeOther, "/member/me")
			} else {
				session.AddFlash("Invalid username or password")
				session.Save()
				c.Redirect(http.StatusSeeOther, "/site/login")
			}

		}
	}

}

//Logout to handle logout request
func Logout(c *gin.Context) {

	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		session.AddFlash("Invalid session token")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/site/index")
	}

	session.Delete(userkey)
	session.Delete(loggedInKey)

	if err := session.Save(); err != nil {
		session.AddFlash("Error at saving session")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/site/index")
	}
	session.AddFlash("Logged out successfully")
	session.Save()
	c.Redirect(http.StatusSeeOther, "/site/login")

}

// RegisterForm shows the form
func RegisterForm(c *gin.Context) {
	session := sessions.Default(c)
	flash := session.Flashes()
	session.Save()
	c.HTML(http.StatusOK, "register.html", gin.H{
		"Title":      "Register",
		"Content":    "Register Here",
		"IsLoggedIn": session.Get(loggedInKey),
		"flashes":    flash,
		"roleID":     session.Get("roleID"),
	})
}

//RegisterHandler handle register request
func RegisterHandler(c *gin.Context) {
	session := sessions.Default(c)
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")
	var id int

	if username == "" || password == "" || email == "" || firstName == "" || lastName == "" {
		session.AddFlash("All fields are required")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/site/register")
	} else {
		db := general.GetDB()
		defer general.CloseDB(db)

		err := db.Raw("SELECT id FROM `users` WHERE username = ? OR email = ?", username, email).Row().Scan(&id)

		if err != nil {

			hasPassword, _ := general.HashPassword(password)
			db.Create(&models.User{
				Username:  username,
				Password:  hasPassword,
				Email:     email,
				FirstName: firstName,
				LastName:  lastName,
				RoleID:    2,
			})
			session.AddFlash("Registered Successfully")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/site/login")

		} else {

			session.AddFlash("Username or email already exists")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/site/register")

		}
	}

}
