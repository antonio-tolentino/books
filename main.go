package main

import (
	"books/storage"
	"books/web"
	"fmt"

	"log"
	"net/http"
)

func main() {

	// Init Storage
	booksRepo := storage.NewBooksMemoryStorage()

	// Init Web Server
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
