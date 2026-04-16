package storage

import (
	"books/models"
	"log"
)

// Mock examples of books
func mockData() map[int]models.Book {

	log.Println("Mocking data...")

	books := make(map[int]models.Book)

	books[448743] = models.Book{Isbn: 448743,
		Title:  "Book One",
		Author: &models.Author{Firstname: "Agatha", Lastname: "Black"}}

	books[556798] = models.Book{Isbn: 556798,
		Title:  "Book Two",
		Author: &models.Author{Firstname: "Steve", Lastname: "White"}}

	books[123456] = models.Book{Isbn: 123456,
		Title:  "Book Three",
		Author: &models.Author{Firstname: "John", Lastname: "Doe"}}

	books[987654] = models.Book{Isbn: 987654,
		Title:  "Book Four",
		Author: &models.Author{Firstname: "Jane", Lastname: "Doe"}}

	return books

}
