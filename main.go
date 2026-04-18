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

	// BUG: unreachable code
	if false {
		fmt.Println("This is unreachable")
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
