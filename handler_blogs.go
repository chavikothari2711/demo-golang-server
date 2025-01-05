package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/chavikothari2711/demo-golang-server/internal/database"
)

func (apiCfg *apiConfig) handleCreateBlogs(w http.ResponseWriter, r *http.Request, existingUser database.User) {

	type parameters struct {
		Title      string `json:"title"`
		Visibility string `json:"visibility"`
		Body       string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "ERROR PARSING JSON")
		log.Println(err)
		return
	}

	visibilityType, err := apiCfg.DB.GetVisibilityId(r.Context(), params.Visibility)
	if err != nil {
		respondWithError(w, 404, "visibilityType is not present, check payload")
		log.Println(visibilityType)
		return
	}

	existingBlog, err := apiCfg.DB.GetBlogByTilte(r.Context(), params.Title)
	if err == nil {
		respondWithError(w, 404, "existingBlog is not present, check payload")
		log.Println(existingBlog)
		return
	}

	blog, err := apiCfg.DB.CreateBlogs(r.Context(), database.CreateBlogsParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Body:       params.Body,
		Title:      params.Title,
		UserID:     existingUser.ID,
		Visibility: visibilityType.ID,
	})

	if err != nil {
		log.Println(err)
		respondWithError(w, 400, "ERROR IN CREATING BLOGS")
		return
	}

	respondWithJSON(w, 200, blog)
}
