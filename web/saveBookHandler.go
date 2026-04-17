package web

import (
	"books/models"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"
)

func (s *StorageHandler) saveBookHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// check if isbn was converted correctly
	if book.Isbn < 100000 {
		slog.Warn("Book not saved", "error", errors.New("Valid isbn is required."), "isbn", book.Isbn)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := s.Storage.Save(ctx, book)
	if err != nil {
		slog.Warn("Book not saved", "error", err, "isbn", book.Isbn)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
