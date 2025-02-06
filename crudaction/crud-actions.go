package crudaction

import (
	"encoding/json"
	"net/http"
	"simplecrud/model"
	"simplecrud/utils"

	"github.com/gorilla/mux"
)

var Animes = model.AnimeList

func GetAnimes(w http.ResponseWriter, r *http.Request) {

	// Set the content type and
	// encode the slice of Animes into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Animes)
}

func GetAnime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Set the content type as "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Check for the ID to be fetched by looping through Animes slice
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
	utils.DeleteRecord(params["id"], &Animes)

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
	if utils.CheckID(Anime.ID, &Animes) {

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
	utils.DeleteRecord(params["id"], &Animes)

	// Populate the updated request into the Anime variable
	json.NewDecoder(r.Body).Decode(&Anime)

	Anime.ID = params["id"]

	// Append the update Anime to the existing slice of Anime
	Animes = append(Animes, Anime)

	// Set the content type and
	// Encode the slice of Animes into "w" i.e. Response Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Animes)
}
