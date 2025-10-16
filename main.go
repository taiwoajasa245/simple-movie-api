package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/simple-movie-api/handlers"
	"github.com/simple-movie-api/models"
)

func main() {
	router := chi.NewRouter()

	// useful middleware
	router.Use(middleware.Logger)    // log requests
	router.Use(middleware.Recoverer) // recover from panic

	// API routes
	router.Route("/movies", func(r chi.Router) {
		r.Get("/", handlers.GetAllMovies)       // GET /movies
		r.Post("/", handlers.CreateMovie)       // POST /movies
		r.Get("/{id}", handlers.GetMovieById)   // GET /movies/{id}
		r.Put("/{id}", handlers.UpdateMovie)    // PUT /movies/{id}
		r.Delete("/{id}", handlers.DeleteMovie) // DELETE /movies/{id}
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		res := models.APIResponse{
			Status:  http.StatusOK,
			Success: true,
			Message: "Movie API is running",
		}
		json.NewEncoder(w).Encode(res)
	})

	// Print fancy banner
	banner := figure.NewFigure("Movie API", "isometric3", true)
	cyan := color.New(color.FgHiCyan).SprintFunc()
	banner.Print()

	fmt.Println(color.HiBlackString("───────────────────────────────"))
	fmt.Printf("🎬  %s\n", cyan("Server running on http://localhost:8080"))
	fmt.Println(color.HiBlackString("───────────────────────────────"))

	// log.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
