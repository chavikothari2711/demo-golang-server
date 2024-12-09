package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/chavikothari2711/demo-golang-server/internal/database"
)

func (apiCfg *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		ID    uuid.UUID `json:id`
		Name  string    `json:name`
		Email string    `json:email`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "ERROR PARSING JSON")
		return
	}
	user, err := apiCfg.DB.UpdateUser(r.Context(), database.UpdateUserParams{
		ID:        params.ID,
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Email:     params.Email,
	})

	if err != nil {
		respondWithError(w, 400, "ERROR IN CREATING USER")
		return
	}

	respondWithJSON(w, 200, databaseUpdateUserToUser(user))
}
