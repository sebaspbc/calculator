package main

import (
	"log"
	"net/http"
	"os"

	"Calculator-backend/internal/handlers"
	"Calculator-backend/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	allowedOrigins := os.Getenv("FRONTEND_ORIGIN")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:5173"
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{allowedOrigins},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	calcService := service.NewCalculatorService()
	calcHandler := handlers.NewCalculatorHandler(calcService)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	r.Post("/api/v1/calculate", calcHandler.Calculate)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port :%s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
