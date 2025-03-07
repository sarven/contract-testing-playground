package main

import (
	"fmt"
	"go-rest-api/internal/server"
	"log"
	"net/http"
)

func main() {
	port := 8080
	router := server.SetupServer()

	log.Printf("Server running on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
