package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/deepaksinghkushwah/module-test1/app/models"
)

// Hello return
func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	users := []models.User{}
	for i := 0; i < 5; i++ {
		user := models.User{ID: i, FirstName: strconv.Itoa(i), LastName: " Hello"}
		users = append(users, user)
	}
	b, error := json.Marshal(users)
	if error != nil {
		log.Fatal(error)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(b)
}
