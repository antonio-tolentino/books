package web

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Get all books
func (s *StorageHandler) listAllBooksHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// get all books
	books := s.Storage.GetAll(ctx)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}
