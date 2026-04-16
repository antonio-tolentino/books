package web

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// List Book Handler
func (s *StorageHandler) listBookHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// get isbn from url
	isbnString := mux.Vars(r)["isbn"]

	// convert to int
	isbn, err := strconv.Atoi(isbnString)
	if err != nil {
		log.Printf("Book not found: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get book
	book, err := s.Storage.GetByIsbn(ctx, isbn)
	if err != nil {
		log.Printf("Book not found: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(book)
}
