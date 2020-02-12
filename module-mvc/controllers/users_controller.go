package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"module-mvc/domain"
	"module-mvc/helpers"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

//GetUser return single user
func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	db := helpers.GetDB()
	defer helpers.CloseDB(db)
	row := db.QueryRow("SELECT id, first_name, last_name FROM `user` WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	var user domain.User
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(200, gin.H{
				"msg": "No record found with id " + fmt.Sprintf("%d", id),
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

//GetAllUserNew return all user
func GetAllUserNew(c *gin.Context) {
	db := helpers.GetDB()
	defer helpers.CloseDB(db)

	limit, err := strconv.ParseInt(c.Param("limit"), 10, 64)
	if err != nil {
		limit = 100
	}
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
