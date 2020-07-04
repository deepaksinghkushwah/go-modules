package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// User struct
type User struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		var users []User
		for i := 1; i < 100; i++ {
			users = append(users, User{Name: "Test " + strconv.Itoa(i), Age: (30 + i)})
		}

		json, _ := json.Marshal(users)
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})

	http.ListenAndServe(":8080", nil)
}
