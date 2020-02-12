package controllers

import (
	"log"
	"os"

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
