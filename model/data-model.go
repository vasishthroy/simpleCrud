package model

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstname"`
	Lastname  string `json: "lastname"`
}

var MovieList = []Movie{
	{ID: "1", Isbn: "23401", Title: "One Piece", Director: &Director{Firstname: "Luffy", Lastname: "Monkey D"}},
	{ID: "2", Isbn: "321451", Title: "Fairy Tail", Director: &Director{Firstname: "Natsu", Lastname: "Dragoneel"}},
}
