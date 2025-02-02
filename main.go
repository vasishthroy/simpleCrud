package main

// "log"
import (
	"fmt"
	"simpleCrud/model"
)

var movies []model.Movie

func main() {

	movies = append(
		movies,
		model.Movie{
			ID:    "1",
			Isbn:  "23401",
			Title: "One Piece",
			Director: &model.Director{
				Firstname: "Luffy",
				Lastname:  "Monkey D",
			},
		},
	)

	movies = append(
		movies,
		model.Movie{
			ID:    "2",
			Isbn:  "11111",
			Title: "One Piece",
			Director: &model.Director{
				Firstname: "Natsu",
				Lastname:  "Dragneel",
			},
		},
	)

	fmt.Printf("Starting server at port 8000\n")

}
