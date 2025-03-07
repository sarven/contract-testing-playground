package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/server"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Given
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := &repository.PostgresUserRepository{Conn: conn}

	user := &model.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	err = repo.CreateUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Failed to add user to database: %v", err)
	}

	// When
	router := server.SetupServer()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/users/%d", user.ID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then
	if http.StatusOK != rr.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, rr.Code)
	}

	var returnedUser model.User
	err = json.NewDecoder(rr.Body).Decode(&returnedUser)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if returnedUser.Name != user.Name || returnedUser.Email != user.Email {
		t.Errorf("Expected user %v. Got %v\n", user, returnedUser)
	}
}
