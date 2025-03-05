package main

import (
    "log"
    "net/http"
    "go-rest-api/internal/handler"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/user/{id}", handler.GetUser).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", r))
}
