package main

import (
	"fmt"
	"log"
	"net/http"

	"simpleCrud/crudaction"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Starting server at port 8000\n")
	r := mux.NewRouter()
	r.HandleFunc("/movies", crudaction.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", crudaction.GetMovie).Methods("GET")

	r.HandleFunc("/movies", crudaction.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", crudaction.UpdateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", crudaction.DeleteMovie).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}
