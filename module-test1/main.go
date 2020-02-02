package main

import (
	"net/http"
	"os"

	"github.com/deepaksinghkushwah/module-test1/app/route"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Applicatio has been started on : 8080")
	log.Info("Access site on : http://localhost:8080/api/v1/hello")
	log.Info("Access site on : http://localhost:8081/api/v2/hello")
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("Site")
	log.Info(key)

	//creating instance of mux
	route1 := route.GetRoutes()
	route2 := route.GetRoutesV2()

	go func() { log.Fatal(http.ListenAndServe(":8080", route1)) }()
	go func() { log.Fatal(http.ListenAndServe(":8081", route2)) }()
	select {}

}
