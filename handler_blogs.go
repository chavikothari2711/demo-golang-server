package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/chavikothari2711/demo-golang-server/internal/database"
)

func (apiCfg *apiConfig) handleCreateBlogs(w http.ResponseWriter, r *http.Request, user database.User) {

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
		respondWithError(w, 404, "blog with same title is present ")
		log.Println(existingBlog)
		return
	}

	blog, err := apiCfg.DB.CreateBlogs(r.Context(), database.CreateBlogsParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Body:       params.Body,
		Title:      params.Title,
		UserID:     user.ID,
		Visibility: visibilityType.ID,
	})

	if err != nil {
		log.Println(err)
		respondWithError(w, 400, "ERROR IN CREATING BLOGS")
		return
	}

	respondWithJSON(w, 200, blog)
}

func (apiCfg *apiConfig) handleGetUserBlogs(w http.ResponseWriter, r *http.Request, user database.User) {
	blogs, err := apiCfg.DB.GetUserBlogs(r.Context(), user.ID)

	if err != nil {
		log.Println(err)
		respondWithError(w, 400, "ERROR IN GETTING BLOGS")
		return
	}

	respondWithJSON(w, 200, blogs)
}

func (apiCfg *apiConfig) handleGetAllBlogs(w http.ResponseWriter, r *http.Request) {

	visibility := r.URL.Query().Get("visibility")
	if visibility == "" {
		respondWithError(w, 400, "Missing 'visibility' query parameter")
		return
	}

	visibilityType, err := apiCfg.DB.GetVisibilityId(r.Context(), visibility)
	if err != nil {
		respondWithError(w, 404, "Visibility type not found")
		log.Println(err)
		return
	}

	blogs, err := apiCfg.DB.GetAllTypeBlogs(r.Context(), visibilityType.ID)
	if err != nil {
		log.Println(err)
		respondWithError(w, 500, "Error fetching blogs")
		return
	}

	respondWithJSON(w, 200, blogs)
}

func (apiCfg *apiConfig) handleDeleteBlog(w http.ResponseWriter, r *http.Request, user database.User) {
	blogId := r.URL.Query().Get("id")
	if blogId == "" {
		respondWithError(w, 400, "Missing 'blogId' query parameter")
		return
	}

	blogUUID, err := uuid.Parse(blogId)
	if err != nil {
		respondWithError(w, 400, "Invalid 'id' format. Must be a valid UUID.")
		return
	}

	existingBlog, err := apiCfg.DB.GetBlog(r.Context(), blogUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find blog")
		return
	}

	_, err = apiCfg.DB.DeleteBlog(r.Context(), database.DeleteBlogParams{
		ID:     existingBlog.ID,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't delete blog")
		return
	}

	respondWithJSON(w, 200, "Delete succesfully")
}

func (apiCfg *apiConfig) handleUpdateBlog(w http.ResponseWriter, r *http.Request, user database.User) {
	blogId := r.URL.Query().Get("id")
	if blogId == "" {
		respondWithError(w, 400, "Missing 'blogId' query parameter")
		return
	}

	blogUUID, err := uuid.Parse(blogId)
	if err != nil {
		respondWithError(w, 400, "Invalid 'id' format. Must be a valid UUID.")
		return
	}

	_, err = apiCfg.DB.GetBlog(r.Context(), blogUUID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't find blog")
		return
	}

	type parameters struct {
		Title      string `json:"title"`
		Visibility string `json:"visibility"`
		Body       string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "ERROR PARSING JSON")
		log.Println(err)
		return
	}

	existingBlog, err := apiCfg.DB.GetBlogByTilte(r.Context(), params.Title)
	if err == nil {
		respondWithError(w, 404, "blog with same title is present ")
		log.Println(existingBlog)
		return
	}

	visibilityType, err := apiCfg.DB.GetVisibilityId(r.Context(), params.Visibility)
	if err != nil {
		respondWithError(w, 404, "visibilityType is not present, check payload")
		log.Println(visibilityType)
		return
	}

	blog, err := apiCfg.DB.UpdateUserBlog(r.Context(), database.UpdateUserBlogParams{
		Body:       params.Body,
		Title:      params.Title,
		Visibility: visibilityType.ID,
		ID:         existingBlog.ID,
		UserID:     user.ID,
	})

	if err != nil {
		log.Println(err)
		respondWithError(w, 400, "ERROR IN UPDATING BLOGS")
		return
	}

	respondWithJSON(w, 200, blog)
}
