package main

import (
	"fmt"
	"net/http"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := "Ivson"
	fmt.Fprintf(w, "Hello, I'm %s.", name)
}
