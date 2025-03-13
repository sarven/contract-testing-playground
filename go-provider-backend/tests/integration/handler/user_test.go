package tests

import (
	"encoding/json"
	"fmt"
	"go-rest-api/internal/model"
	"go-rest-api/internal/server"
	"go-rest-api/tests/fixtures"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	// Given
	user := fixtures.GivenUser("John Doe", "john.doe@example.com")

	// When
	router := server.SetupRouter()
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/users/%d", user.ID), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Then
	if http.StatusOK != rr.Code {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, rr.Code)
	}

	var returnedUser model.User
	err := json.NewDecoder(rr.Body).Decode(&returnedUser)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if returnedUser.Name != "John Doe" || returnedUser.Email != "john.doe@example.com" {
		t.Errorf("Expected user %v. Got %v\n", user, returnedUser)
	}
}
