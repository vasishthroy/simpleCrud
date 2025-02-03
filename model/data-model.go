package model

import (
	"math/rand"
	"strconv"
)

type Anime struct {
	ID            string     `json:"id"`
	Isbn          string     `json:"isbn"`
	Title         string     `json:"title"`
	MainCharacter *Character `json:"maincharacter"`
}

type Character struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func CreateID() string {
	return strconv.Itoa(rand.Intn(1000))
}

var AnimeList = []Anime{
	{ID: CreateID(), Isbn: "23401", Title: "One Piece", MainCharacter: &Character{FirstName: "Luffy", LastName: "Monkey D"}},
	{ID: CreateID(), Isbn: "321451", Title: "Fairy Tail", MainCharacter: &Character{FirstName: "Natsu", LastName: "Dragoneel"}},
}

var TotalAnime = len(AnimeList)
