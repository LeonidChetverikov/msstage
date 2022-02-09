package main

import (
	"fmt"

	"net/http"
)

func main() {
	// The "HandleFunc" method accepts a path and a function as arguments
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello MSStage!")
}
