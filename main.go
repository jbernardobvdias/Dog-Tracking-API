package main

import (
	"dog-tracking/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/dogs", handler.GetAllDogs)

	http.ListenAndServe(":8080", nil)
}
