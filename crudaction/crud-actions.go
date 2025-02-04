package crudaction

import (
	"encoding/json"
	"net/http"
	"simpleCrud/model"
	"simpleCrud/utils"

	"github.com/gorilla/mux"
)

var Animes = model.AnimeList

func GetAnimes(w http.ResponseWriter, r *http.Request) {
	// // Print the response body
	// fmt.Printf("Request body %v\n", r.Body)

	// Set the content type and
	// encode the slice of Animes into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Animes)
}

func GetAnime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Set the content type as "application/json"
	w.Header().Set("Content-Type", "application/json")
	for _, item := range Animes {
		if item.ID == params["id"] {
			// Encode the slice of Animes into "w" i.e. Response Writer
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func DeleteAnime(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	// Check for the ID that needs to be deleted
	for index, item := range Animes {
		if item.ID == params["id"] {
			// Delete the record from the slice
			Animes = append(Animes[:index], Animes[index+1:]...)
			break
		}
	}

	// Set the content type and
	// encode the slice of Animes into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Animes)
}

func CreateAnime(w http.ResponseWriter, r *http.Request) {

	var Anime model.Anime

	// Populate the new Anime record(request) into the Anime variable
	json.NewDecoder(r.Body).Decode(&Anime)

	// Check if ID provided in Request body is used
	if utils.CheckID(Anime.ID, Animes) {

		// Create an ID for the new Anime that's been added
		Anime.ID = utils.CreateID()
	}

	Animes = append(Animes, Anime)

	// Set the content type and
	// encode the slice of Animes into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Animes)
}

func UpdateAnime(w http.ResponseWriter, r *http.Request) {

	var Anime model.Anime

	params := mux.Vars(r)

	// Check for the Anime ID that needs to be updated
	for index, item := range Animes {
		if item.ID == params["id"] {
			// Delete the record from the slice
			Animes = append(Animes[:index], Animes[index+1:]...)
			break
		}
	}

	// Populate the updated request into the Anime variable
	json.NewDecoder(r.Body).Decode(&Anime)
	Anime.ID = params["id"]
	// Append the update Anime to the existing slice of Anime
	Animes = append(Animes, Anime)

	// Set the content type and
	// encode the slice of Animes into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Animes)
}
