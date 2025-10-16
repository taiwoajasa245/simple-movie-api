package db

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/simple-movie-api/models"
)

var (
	mu sync.Mutex

	movies = []models.Movie{}
)

// create new id
func newID() string {

	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func init() {
	movies = []models.Movie{
		{ID: newID(), Title: "Inception", Director: "Christopher Nolan", Year: 2010, Genre: "Sci-Fi"},
		{ID: newID(), Title: "The Matrix", Director: "Lana Wachowski, Lilly Wachowski", Year: 1999, Genre: "Sci-Fi"},
	}
}

func GetAllMovies() []models.Movie {
	mu.Lock()
	defer mu.Unlock()

	// return the copy to prevent caller mutation internal slice
	movieCopy := make([]models.Movie, len(movies))

	// copies movies into moviesCopy
	copy(movieCopy, movies)

	return movieCopy
}

func GetMovieById(id string) (models.Movie, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, value := range movies {

		if value.ID == id {
			return value, nil
		}
	}

	return models.Movie{}, errors.New("movie not found")

}

func CreateMovie(movie models.Movie) models.Movie {
	mu.Lock()
	defer mu.Unlock()

	movie.ID = newID()
	movies = append(movies, movie)
	return movie
}

func UpdateMovie(m models.Movie, id string) (models.Movie, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, movie := range movies {
		if movie.ID == id {
			m.ID = id

			movies[i] = m
			return movies[i], nil
		}
	}
	return models.Movie{}, errors.New("movie not found")
}

func DeleteMovie(id string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			return nil
		}
	}
	return errors.New("movie not found")
}
