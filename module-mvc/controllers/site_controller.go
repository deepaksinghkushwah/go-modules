package controllers

import (
	"module-mvc/helpers"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Index function
func Index(c *gin.Context) {
	p := struct {
		PageTitle string
		Content   string
	}{
		PageTitle: "Index Page",
		Content:   "This is sample content",
	}
	c.HTML(200, "index.html", p)
}

func PopupDB(c *gin.Context) {
	db := helpers.GetDB()

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
