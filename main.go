package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	serve(port)
}

func serve(port string) error {
	http.HandleFunc("/", routeMatch)
	return http.ListenAndServe(":"+port, nil)
}

func routeMatch(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "hello world", http.StatusOK)
}
