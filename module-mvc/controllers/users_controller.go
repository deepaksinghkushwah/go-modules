package controllers

import (
	"database/sql"
	"log"
	"module-mvc/helpers"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

//User struct
type User struct {
	UserID    int64          `json:"user_id"`
	FirstName sql.NullString `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
}

//GetUser return single user
func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	db := helpers.GetDB()
	row := db.QueryRow("SELECT id, first_name, last_name FROM `user` WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	var user User
	err = row.Scan(&user.UserID, &user.FirstName, &user.LastName)
	if err != nil {
		panic(err)
	}

	c.JSON(200, user)

}

//GetAllUserNew return all user
func GetAllUserNew(c *gin.Context) {
	db := helpers.GetDB()

	limit, err := strconv.ParseInt(c.Param("limit"), 10, 64)
	if err != nil {
		limit = 100
	}
	users := []User{}
	//runtime.GOMAXPROCS(4)
	ch := make(chan User)
	quit := make(chan int)
	defer helpers.CloseDB(db)
	go func(ch chan User) {
		rows, err := db.Query("SELECT id, first_name, last_name FROM `user` LIMIT 0,?", limit)

		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var tUser User
			err = rows.Scan(&tUser.UserID, &tUser.FirstName, &tUser.LastName)
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
