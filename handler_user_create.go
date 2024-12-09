package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/chavikothari2711/demo-golang-server/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name  string `json:name`
		Email string `json:email`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "ERROR PARSING JSON")
		return
	}

	existingUser, err := apiCfg.DB.GetUser(r.Context(), params.Email)
	if err == nil {
		respondWithError(w, 404, "User already present with same email")
		log.Println(existingUser)
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
	})

	if err != nil {
		respondWithError(w, 400, "ERROR IN CREATING USER")
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}