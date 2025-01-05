package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/chavikothari2711/demo-golang-server/internal/database"
)

func (apiCfg *apiConfig) handlerCreateVisibilityType(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		BlogType string `json:"blogType"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "ERROR PARSING JSON")
		return
	}

	existingType, err := apiCfg.DB.GetVisibilityId(r.Context(), params.BlogType)
	if err == nil {
		respondWithError(w, 404, "Visibilty Type already present with same name")
		log.Println(existingType)
		return
	}

	visibility, err := apiCfg.DB.CreateBlogVisibilityType(r.Context(), database.CreateBlogVisibilityTypeParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Visibilitytype: params.BlogType,
	})

	if err != nil {
		respondWithError(w, 400, "ERROR IN CREATING Visibility")
		return
	}

	respondWithJSON(w, 200, visibility)
}

func (apiCfg *apiConfig) handlerGetAllVisibilityType(w http.ResponseWriter, r *http.Request) {
	existingType, err := apiCfg.DB.GetAllVisibilityType(r.Context())
	if err != nil {
		respondWithError(w, 404, "Visibilty Type not present")
		log.Println(existingType)
		return
	}
	respondWithJSON(w, http.StatusOK, existingType)
}
