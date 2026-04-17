package main

import (
	"books/storage"
	"books/web"
	"log/slog"
	"net/http"
	"os"
)

func main() {

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
