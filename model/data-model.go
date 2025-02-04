package model

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

var AnimeList []Anime
