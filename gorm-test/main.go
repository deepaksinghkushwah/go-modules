package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Product struct
type Product struct {
	gorm.Model
	Code   string
	Price  uint
	Status uint
}

func main() {
	db, err := gorm.Open("mysql", "root:Deepak123@@/test?charset=utf8&parseTime=true&loc=Local") // for mysql
	//db, err := gorm.Open("sqlite3", "./test.db") // for sqlite
	checkError(err)

	defer db.Close()

	// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{Code: "L121213", Price: 2000, Status: 1})

	// read
	product := Product{}
	db.First(&product, 1)

	// update
	db.Model(&product).Update("Price", 11000)

	// delete
	//db.Delete(&product)

}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
