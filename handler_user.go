package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bhattrahul525/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		jsonError(w, 400, "Error parsing JSON ")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		jsonError(w, 400, fmt.Sprintf("Couldn't create user: ", err))
		return
	}

	jsonResponse(w, 200, user)
}
