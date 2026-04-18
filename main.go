package main

import (
	"books/storage"
	"books/web"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// main is the entry point of the books API.
// It initializes the logger, storage and web server.
// It then starts the web server using slog.
func main() {

	x := 10
	if x > 5 {
		if x > 6 {
			if x > 7 {
				fmt.Println("High complexity")
			}
		}
	}

	// Init Logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Init Storage
	booksRepo := storage.NewBooksMemoryStorage()

	// Init Web Server
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	slog.Info("Server running at http://localhost:8080")

	// Start Server using slog
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error(err.Error())
	}
}
