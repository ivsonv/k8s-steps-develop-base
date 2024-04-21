package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Healthz(w http.ResponseWriter, r *http.Request) {

	duration := time.Since(startedAt)
	if duration.Seconds() < 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("SERVE_APP_NAME")
	fmt.Fprintf(w, "Hello, I'm %s.", name)
}
