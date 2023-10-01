package main

import (
	"fmt"
	"net/http"
)

const port = ":8887"

func main() {

	http.HandleFunc("/hello", Hello)

	fmt.Println("listening: " + port)
	http.ListenAndServe(port, nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
