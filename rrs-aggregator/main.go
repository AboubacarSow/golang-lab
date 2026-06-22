package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Printf("Unable to load .env file")
		os.Exit(0)
	}

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
	v1_router.Get("/error",errorHandler)

	router.Mount("/api/v1", v1_router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Printf("Server starting on port :%s\n", port)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error occured while starting server")
	}
}
