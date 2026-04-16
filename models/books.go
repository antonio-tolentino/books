package models

// Author
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Book
type Book struct {
	Isbn   int     `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

func (B Book) String() string {
	return B.Title + " by " + B.Author.Firstname
}
