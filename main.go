package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	podName, _ := os.Hostname()
	greeting := os.Getenv("GREETING")
	fmt.Fprintf(w, "%s %s", greeting, podName)
}
