package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Product struct
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "root:deepak@/test?parseTime=true")
	checkError(err)

	defer db.Close()

	// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
	db.AutoMigrate(&Product{})

	// create
	//db.Create(&Product{Code: "L121213", Price: 1000})

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
