package crudaction

import (
	"encoding/json"
	"fmt"
	"net/http"
	"simpleCrud/model"

	"github.com/gorilla/mux"
)

var movies = model.MovieList

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request body %v\n", r.Body)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}
