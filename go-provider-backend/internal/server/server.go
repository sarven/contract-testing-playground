package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"go-rest-api/internal/handler"
	"go-rest-api/internal/repository"
)

func SetupServer() http.Handler {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	userRepo := &repository.PostgresUserRepository{Conn: conn}

	userHandler := handler.NewUserHandler(userRepo)

	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	return r
}
