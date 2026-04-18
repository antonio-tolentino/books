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

func (b Book) String() string {
	return b.Title + " by " + b.Author.Firstname
}
