package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/chavikothari2711/demo-golang-server/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT not available")
		return
	}

	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("DB URL not available")
		return
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error in connecting sql database: ", err)
		return
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	//health routes
	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/error", handlerError)
	// users routes
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Put("/users", apiCfg.middlewareAuth(apiCfg.handlerUpdateUser))
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	// visibilityType routes
	v1Router.Post("/visibility", apiCfg.handlerCreateVisibilityType)
	v1Router.Get("/visibility-types", apiCfg.handlerGetAllVisibilityType)
	// blog routes
	v1Router.Post("/blogs", apiCfg.middlewareAuth(apiCfg.handleCreateBlogs))
	v1Router.Get("/blogs/user", apiCfg.middlewareAuth(apiCfg.handleGetUserBlogs))
	v1Router.Get("/blogs", apiCfg.handleGetAllBlogs)
	v1Router.Delete("/blogs", apiCfg.middlewareAuth(apiCfg.handleDeleteBlog))
	v1Router.Put("/blogs", apiCfg.middlewareAuth(apiCfg.handleUpdateBlog))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting on PORT: %s", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
