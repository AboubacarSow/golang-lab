package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AboubacarSow/golang-lab/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Unable to load .env file")
	}
	connectionString := os.Getenv("CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("Unable to get CONNECTION_STRING")
	}

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Could not connect to database", err)
	}

	apiConfig := ApiConfig{
		DB: database.New(conn),
	}
	go scraperSpinner(apiConfig.DB, 3, 10*time.Second)

	router := chi.NewRouter()

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https:\\*", "http:\\*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
	router.Use(corsMiddleware)

	v1_router := chi.NewRouter()

	v1_router.Get("/health", healthHandler)
	v1_router.Get("/error", errorHandler)

	// USER REST API ENDPOINTS
	v1_router.Post("/users", apiConfig.createUserHandler)
	v1_router.Get("/users", apiConfig.authMiddleware(getUserHandler))
	v1_router.Delete("/users/{id}", apiConfig.deleteUserHandler)

	//Feed REST API ENDPOINTS
	v1_router.Post("/feeds", apiConfig.authMiddleware(apiConfig.createFeedHandler))
	v1_router.Get("/feeds", apiConfig.getAllFeedsHandler)
	v1_router.Delete("/feeds/{id}", apiConfig.authMiddleware(apiConfig.DeleteOneFeedHandler))

	// Feed Follow REST API ENDPOINTS
	v1_router.Post("/feed_follows", apiConfig.authMiddleware(apiConfig.createFeedFollowsHandler))
	v1_router.Get("/feed_follows", apiConfig.authMiddleware(apiConfig.getAllFeedFollowsHandler))

	// Test
	//v1_router.Get("/display-rss", outputRssFeedHandler)

	router.Mount("/api/v1", v1_router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server starting on port :%s\n", port)
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal("Error occured while starting server")
	}
}
