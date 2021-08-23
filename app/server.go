package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Congratulations! Your Go application has been successfully deployed on Kubernetes.</h1>")
	if err != nil {
		return
	}
}

func check(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "<h1>Health check</h1>")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/health", check)
	fmt.Println("Server starting...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		return
	}
}
