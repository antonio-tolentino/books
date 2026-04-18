package main

import (
	"books/storage"
	"books/web"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

// FAILURE 3: High Code Duplication
// Standard gates fail if > 3.0% of lines are duplicated.
func triggerDuplicationA() {
	fmt.Println("This block is identical to the one below.")
	fmt.Println("SonarQube will flag this as a duplication issue.")
	items := []string{"one", "two", "three", "four", "five"}
	for _, item := range items {
		fmt.Printf("Processing item: %s\n", item)
	}
}

func triggerDuplicationB() {
	fmt.Println("This block is identical to the one below.")
	fmt.Println("SonarQube will flag this as a duplication issue.")
	items := []string{"one", "two", "three", "four", "five"}
	for _, item := range items {
		fmt.Printf("Processing item: %s\n", item)
	}
}

// main is the entry point of the books API.
// It initializes the logger, storage and web server.
// It then starts the web server using slog.
func main() {

	// FAILURE 1: Security Hotspot - Hardcoded password/secret
	dbPassword := "super_secret_password_123"
	fmt.Println("Connecting to DB with:", dbPassword)

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
