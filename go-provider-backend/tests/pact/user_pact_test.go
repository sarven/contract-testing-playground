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

	brokerURL := os.Getenv("PACT_BROKER_URL")
	brokerToken := os.Getenv("PACT_BROKER_TOKEN")
	providerVersion := os.Getenv("PROVIDER_VERSION")

	if brokerURL == "" || brokerToken == "" || providerVersion == "" {
		t.Fatal("PACT_BROKER_URL, PACT_BROKER_TOKEN, or PROVIDER_VERSION is not set")
	}

	_, err = pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://127.0.0.1:8081",
		BrokerURL:                  brokerURL,
		BrokerToken:                brokerToken,
		PublishVerificationResults: true,
		ProviderVersion:            providerVersion,
	})

	if err != nil {
		t.Fatal(err)
	}
}
