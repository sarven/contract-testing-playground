package tests

import (
	"fmt"
	"go-rest-api/internal/server"
	"go-rest-api/tests/fixtures"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func startProvider() {
	server.SetupServer(8081)
}

func TestServerPact_Verification(t *testing.T) {
	fixtures.GivenUser("john.doe@example.com")

	go startProvider()

	var dir, _ = os.Getwd()
	var pactDir = fmt.Sprintf("%s/../../pacts", dir)
	var logDir = fmt.Sprintf("%s/log", dir)

	pact := &dsl.Pact{
		Provider:                 "Backend",
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

	_, err := pact.VerifyProvider(t, types.VerifyRequest{
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
