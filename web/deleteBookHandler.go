package web

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *StorageHandler) deleteBookHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// get isbn from url
	isbnString := mux.Vars(r)["isbn"]

	// convert to int
	isbn, err := strconv.Atoi(isbnString)
	if err != nil {
		slog.Warn("Book isbn is not valid", "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = s.Storage.Delete(ctx, isbn)
	if err != nil {
		slog.Warn("Book not deleted", "isbn", isbn, "error", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
