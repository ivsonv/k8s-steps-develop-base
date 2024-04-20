package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := "Ivson"
	fmt.Fprintf(w, "Hello, I'm %s.", name)
}
