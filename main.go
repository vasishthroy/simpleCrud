package main

import (
	"fmt"
	"log"
	"net/http"
	"simpleCrud/crudaction"
	"simpleCrud/model"
	"simpleCrud/utils"

	"github.com/gorilla/mux"
)

func main() {

	crudaction.Animes = append(
		crudaction.Animes,
		model.Anime{
			ID:    utils.CreateID(),
			Isbn:  "23401",
			Title: "One Piece",
			MainCharacter: &model.Character{
				FirstName: "Luffy",
				LastName:  "Monkey D",
			},
		},
	)
	crudaction.Animes = append(
		crudaction.Animes,
		model.Anime{
			ID:    utils.CreateID(),
			Isbn:  "321451",
			Title: "Fairy Tail",
			MainCharacter: &model.Character{
				FirstName: "Natsu",
				LastName:  "Dragoneel",
			},
		},
	)

	portNumber := "8080"
	listeningPort := fmt.Sprintf(":%v", portNumber)

	fmt.Printf("Starting server at port %v\n", portNumber)

	r := mux.NewRouter()
	r.HandleFunc("/animes", crudaction.GetAnimes).Methods("GET")
	r.HandleFunc("/animes/{id}", crudaction.GetAnime).Methods("GET")

	r.HandleFunc("/animes", crudaction.CreateAnime).Methods("POST")
	r.HandleFunc("/animes/{id}", crudaction.UpdateAnime).Methods("PUT")

	r.HandleFunc("/animes/{id}", crudaction.DeleteAnime).Methods("DELETE")

	log.Fatal(http.ListenAndServe(listeningPort, r))
}
