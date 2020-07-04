package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/page/{alias}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		alias := vars["alias"]
		fmt.Fprintf(w, "Hello %s", alias)
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You request %s", r.URL.Path)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("static", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", r)
}
