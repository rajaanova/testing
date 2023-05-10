package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil)

}
