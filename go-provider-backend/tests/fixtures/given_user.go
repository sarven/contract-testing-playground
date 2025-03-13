package fixtures

import (
	"context"
	"fmt"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func GivenUser(name string, email string) *model.User {
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
		Name:  name,
		Email: email,
	}
	err = repo.CreateUser(context.Background(), user)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to add user to database: %v", err)
	}

	return user
}
