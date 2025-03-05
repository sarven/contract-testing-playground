package pact

import (
    "testing"
    "github.com/pact-foundation/pact-go/v2/dsl"
    "github.com/pact-foundation/pact-go/v2/models"
)

func TestPact(t *testing.T) {
    pact := dsl.Pact{
        Consumer: "UserConsumer",
        Provider: "UserProvider",
    }

    defer pact.Teardown()

    pact.AddInteraction().
        Given("User with ID 1 exists").
        UponReceiving("A request for user details").
        WithRequest(dsl.Request{
            Method: "GET",
            Path:   dsl.String("/user/1"),
        }).
        WillRespondWith(dsl.Response{
            Status:  200,
            Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
            Body: dsl.Match(&models.User{
                ID:    1,
                Name:  "John Doe",
                Email: "john.doe@example.com",
            }),
        })

    err := pact.Verify(func() error {
        // Call the actual service here
        return nil
    })
    if err != nil {
        t.Fatalf("Error on Verify: %v", err)
    }
}
