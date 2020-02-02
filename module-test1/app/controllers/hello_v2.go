package controllers

import (
	"fmt"
	"net/http"
)

// HelloV2 return
func HelloV2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "From v2")
}
