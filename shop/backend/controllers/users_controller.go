package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"shop/domain"
	"shop/helpers"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var wg sync.WaitGroup

//GetUser return single user
func GetUser(c *gin.Context) {
	id := c.Query("id")
	var user domain.User
	db := helpers.GetDB()
	defer helpers.CloseDB(db)
	err := db.QueryRow("SELECT id, first_name, last_name FROM `user` WHERE id = ?", id).Scan(&user.ID, &user.FirstName, &user.LastName)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{
				"msg": "No record found with id " + id,
			})
			return
		}
		c.JSON(200, gin.H{
			"msg": err.Error(),
		})
		return

	}

	c.JSON(200, user)

}

//GetAllUsers return all user
func GetAllUsers(c *gin.Context) {
	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	limit := c.DefaultQuery("limit", "100")
	users := []domain.User{}
	//runtime.GOMAXPROCS(4)
	ch := make(chan domain.User)
	quit := make(chan int)

	go func(ch chan domain.User) {
		rows, err := db.Query("SELECT id, first_name, last_name FROM `user` LIMIT 0,?", limit)

		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var tUser domain.User
			err = rows.Scan(&tUser.ID, &tUser.FirstName, &tUser.LastName)
			if err != nil {
				log.Fatal(err)
			} else {
				ch <- tUser
			}
		}
		/*for i = 1; i <= limit; i++ {
			stri := strconv.Itoa(int(i))
			ch <- User{UserID: i, FirstName: "FirstName " + stri, LastName: "LastName " + stri}
		}*/
		defer rows.Close()
		quit <- 0
	}(ch)
	for {
		select {
		case u, ok := <-ch:
			if ok {
				users = append(users, u)
			}

		case <-quit:
			c.JSON(200, users)
			return
		}
	}

}

//LoginHandler to handle login request
func LoginHandler(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	//log.Println("username", username, "password:", password)

	if strings.Trim("username", " ") == "" || strings.Trim("password", " ") == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "You must provide username and password",
			"status": "0",
		})
		return
	}
	db := helpers.GetDB()
	defer helpers.CloseDB(db)
	row := db.QueryRow("SELECT id, username, password FROM `user` WHERE `username` = ?", username)

	var user domain.User

	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "No user found",
			"status": "0",
		})
		return
	}
	if helpers.CheckPasswordHash(password, user.Password) {
		uuid := uuid.NewV4()
		var accessToken string = uuid.String()
		accessTokenExpireDate := time.Now().Add(time.Hour * 24 * 60).Format("2006-01-02")
		_, err = db.Exec("UPDATE `user` SET access_token = ?, access_token_expire_date = ? WHERE id = ?", accessToken, accessTokenExpireDate, user.ID)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "Login Successful",
				"status": "1",
				"token":  accessToken,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg":    "Error at updating access token",
			"status": "0",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"msg":    "Invalid username or password",
		"status": "0",
	})
	return
}

// there is no need for logout handler at backend because we are using localStorage at frontend vue app
// so we can comment or ignore this function, it was created because we were

//RegisterHandler handle register request
func RegisterHandler(c *gin.Context) {

	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	firstName := c.PostForm("firstname")
	lastName := c.PostForm("lastname")
	var id int

	if username == "" || password == "" || email == "" || firstName == "" || lastName == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "All fields are required",
			"status": "0",
		})
		return
	}

	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	row := db.QueryRow("SELECT id FROM `user` WHERE username = ? OR email = ?", username, email)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			hasPassword, _ := helpers.HashPassword(password)
			_, err = db.Exec("INSERT INTO `user` SET `username` = ?, password = ?, email = ?, first_name = ?, last_name =?", username, hasPassword, email, firstName, lastName)

			c.JSON(http.StatusOK, gin.H{
				"msg":    "Registered Successfully, please login",
				"status": "1",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":    err.Error(),
			"status": "0",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "Username or email already exists",
		"status": "0",
	})
	return

}
