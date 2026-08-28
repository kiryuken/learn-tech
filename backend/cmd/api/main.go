package main

import (
	"log"
	"net/http"
	"time"

	"learn-tech/backend/internal/api"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           api.New(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("listening on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}
