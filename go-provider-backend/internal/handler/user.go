package handler

import (
    "encoding/json"
    "net/http"
    "strconv"
    "go-rest-api/internal/repository"
    "github.com/gorilla/mux"
)

var repo repository.UserRepository

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := repo.GetUserByID(r.Context(), id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
