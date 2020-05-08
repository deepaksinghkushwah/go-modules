package controllers

import (
	"fmt"
	"log"
	"mvc-gorm/helpers/general"
	"mvc-gorm/models"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var wg sync.WaitGroup

//GetUser return single user
func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	db := general.GetDB()
	defer general.CloseDB(db)
	var user models.User
	err = db.Raw("SELECT id, first_name, last_name FROM `users` WHERE id = ?", id).Row().Scan(&user)
	if err != nil {
		panic(err)
	}

	if user.ID <= 0 {
		c.JSON(200, gin.H{
			"msg": "No record found with id " + fmt.Sprintf("%d", id),
		})
	} else {
		c.JSON(200, user)
	}

}

//GetAllUserNew return all user
func GetAllUserNew(c *gin.Context) {
	db := general.GetDB()
	defer general.CloseDB(db)

	limit, err := strconv.ParseInt(c.Param("limit"), 10, 64)
	if err != nil {
		limit = 100
	}
	users := []models.User{}
	//runtime.GOMAXPROCS(4)
	ch := make(chan models.User)
	quit := make(chan int)

	go func(ch chan models.User) {
		rows, err := db.Raw("SELECT id, first_name, last_name FROM `users` LIMIT 0,?", limit).Rows()

		if err != nil {
			panic(err)
		}
		for rows.Next() {
			var tUser models.User
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
