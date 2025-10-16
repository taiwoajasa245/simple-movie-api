package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/simple-movie-api/db"
	"github.com/simple-movie-api/models"
)

// respondJSON helper
func respondJSON(w http.ResponseWriter, code int, success bool, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := models.APIResponse{
		Status:  code,
		Success: success,
		Message: message,
		Data:    payload,
		// Data: models.Movies{
		// 	Movies: payload,
		// },
	}
	json.NewEncoder(w).Encode(response)
}

// respondError helper
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, false, message, map[string]string{"error": message})
}

// GetMovies -> GET /movies
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies := db.GetAllMovies()

	respondJSON(w, http.StatusOK, true, "Successful", movies)
}

// GetMovie -> GET /movies/{id}
func GetMovieById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	movies, err := db.GetMovieById(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "movie not found")
		return
	}

	respondJSON(w, http.StatusOK, true, "Successful", movies)
}

// CreateMovie -> POST /movies
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	var payload models.Movie

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if payload.Title == "" {
		respondError(w, http.StatusBadRequest, "title is required")
		return
	}

	created := db.CreateMovie(payload)

	respondJSON(w, http.StatusOK, true, "Successful", created)
}

// UpdateMovie -> PUT /movies/{id}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var payload models.Movie

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON body")
		return
	}

	if payload.Title == "" {
		respondError(w, http.StatusBadRequest, "title is required")
		return
	}

	updated, err := db.UpdateMovie(payload, id)

	if err != nil {
		respondError(w, http.StatusNotFound, "movie not found")
		return
	}

	respondJSON(w, http.StatusOK, true, "Successful", updated)
}

// DeleteMovie -> DELETE /movies/{id}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := db.DeleteMovie(id); err != nil {
		respondError(w, http.StatusNotFound, "movie not found")
		return
	}

	respondJSON(w, http.StatusOK, true, "Movie deleted successfully", "")
	w.WriteHeader(http.StatusNoContent) // 204
}
