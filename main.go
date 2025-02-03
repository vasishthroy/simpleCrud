package main

import (
	"fmt"
	"log"
	"net/http"
	"simpleCrud/crudaction"

	"github.com/gorilla/mux"
)

func main() {
	portNumber := "8080"
	fmt.Printf("Starting server at port %v\n", portNumber)

	r := mux.NewRouter()
	r.HandleFunc("/animes", crudaction.GetAnimes).Methods("GET")
	r.HandleFunc("/animes/{id}", crudaction.GetAnime).Methods("GET")

	r.HandleFunc("/animes", crudaction.CreateAnime).Methods("POST")
	r.HandleFunc("/animes/{id}", crudaction.UpdateAnime).Methods("PUT")

	r.HandleFunc("/animes/{id}", crudaction.DeleteAnime).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", portNumber), r))
}
