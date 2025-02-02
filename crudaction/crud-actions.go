package crudaction

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"simpleCrud/model"
	"strconv"

	"github.com/gorilla/mux"
)

var movies = model.MovieList

func GetMovies(w http.ResponseWriter, r *http.Request) {
	// Print the response body
	fmt.Printf("Request body %v\n", r.Body)

	// Set the content type and
	// encode the slice of movies into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Set the content type as "application/json"
	w.Header().Set("Content-Type", "application/json")
	for _, item := range movies {
		if item.ID == params["id"] {
			// Encode the slice of movies into "w" i.e. Response Writer
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// Check for the ID that needs to be deleted
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the record from the slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Set the content type and
	// encode the slice of movies into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	var movie model.Movie

	// Populate the new movie record(request) into the movie variable
	json.NewDecoder(r.Body).Decode(&movie)

	// Create an ID for the new Movie that's been added
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)

	// Set the content type and
	// encode the slice of movies into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

	var movie model.Movie

	params := mux.Vars(r)

	// Check for the Movie ID that needs to be updated
	for index, item := range movies {
		if item.ID == params["id"] {
			// Delete the record from the slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Populate the updated request into the movie variable
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	// Append the update movie to the existing slice of movie
	movies = append(movies, movie)

	// Set the content type and
	// encode the slice of movies into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
