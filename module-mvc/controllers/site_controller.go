package controllers

import (
	"database/sql"
	"log"
	"module-mvc/domain"
	"module-mvc/helpers"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	c.HTML(200, "index.html", gin.H{
		"Title":      "Index Page",
		"Content":    "This is sample content",
		"IsLoggedIn": session.Get(loggedInKey),
		"flashes":    flash,
	})

}

//PopupDB to populate db
func PopupDB(c *gin.Context) {
	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	stmt, err := db.Prepare("INSERT INTO `user` SET username = ?, password = ?, first_name = ?, last_name = ?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 0; i < 50000; i++ {
		stri := strconv.Itoa(i)
		_, err = stmt.Exec("test"+stri, "123456", "test", stri)
		if err != nil {
			panic(err)
		}
	}
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
		db := helpers.GetDB()
		defer helpers.CloseDB(db)
		row := db.QueryRow("SELECT id, username, password FROM `user` WHERE `username` = ?", username)

		var user domain.User
		err := row.Scan(&user.ID, &user.FirstName, &user.Password)
		if err != nil {
			session.AddFlash("User not found")
			session.Save()
			c.Redirect(http.StatusMovedPermanently, "/site/login")
		} else {
			if helpers.CheckPasswordHash(password, user.Password) {
				session.Set(userkey, username)
				session.Set(loggedInKey, true)
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

	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT id FROM `user` WHERE username = ? OR email = ?", username, email)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			hasPassword, _ := helpers.HashPassword(password)
			_, err = db.Exec("INSERT INTO `user` SET `username` = ?, password = ?, email = ?, first_name = ?, last_name =?", username, hasPassword, email, firstName, lastName)

			session.AddFlash("Registered Successfully")
			session.Save()
			c.Redirect(http.StatusSeeOther, "/site/login")
		} else {
			session.AddFlash(err.Error())
			session.Save()
			c.Redirect(http.StatusSeeOther, "/site/register")
		}

	} else {

		session.AddFlash("Username or email already exists")
		session.Save()
		c.Redirect(http.StatusSeeOther, "/site/register")

	}

}
