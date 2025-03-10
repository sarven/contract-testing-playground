package tests

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
	"go-rest-api/internal/server"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func startProvider() {
	server.SetupServer(8081)
}

func TestServerPact_Verification(t *testing.T) {
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

	go startProvider()

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../../pacts", dir)
	var logDir = fmt.Sprintf("%s/log", dir)

	pact := &dsl.Pact{
		Provider:                 "Go Backend Provider",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		DisableToolValidityCheck: true,
	}

	// _, err := pact.VerifyProvider(t, types.VerifyRequest{
	// 	ProviderBaseURL:            "http://127.0.0.1:8080",   //provider's URL
	// 	BrokerURL:                  "https://pen.pactflow.io", //link to your remote Contract broker
	// 	BrokerToken:                "jEQnxw7xWgYRv-3-G7Cx-g",  //your PactFlow token
	// 	PublishVerificationResults: true,
	// 	ProviderVersion:            "1.0.0",
	// })

	// if err != nil {
	// 	t.Fatal(err)
	// }

	_, err = pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8081",
		PactURLs: []string{
			filepath.ToSlash(fmt.Sprintf("%s/PHPBackendConsumer-Backend.json", pactDir)),
			filepath.ToSlash(fmt.Sprintf("%s/ReactFrontend-Backend.json", pactDir)),
		},
	})

	if err != nil {
		t.Fatal(err)
	}
}
