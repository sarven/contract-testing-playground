package tests

import (
//     "context"
    "encoding/json"
    "net/http"
    "testing"
    "go-rest-api/internal/model"
//     "github.com/jackc/pgx/v5"
//     "log"
)

func TestGetUser(t *testing.T) {
    // Create an HTTP client and make a request to the endpoint
    client := &http.Client{}
    req, err := http.NewRequest("GET", "http://localhost:8080/user/1", nil)
    if err != nil {
        t.Fatal(err)
    }

    resp, err := client.Do(req)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // Check the status code
    if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body
    var responseUser model.User
    err = json.NewDecoder(resp.Body).Decode(&responseUser)
    if err != nil {
        t.Fatal(err)
    }

    if responseUser != *user {
        t.Errorf("handler returned unexpected body: got %v want %v", responseUser, *user)
    }
}
