package web

import (
	"context"
	"log"
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
		log.Printf("Book isbn is not valid: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = s.Storage.Delete(ctx, isbn)
	if err != nil {
		log.Printf("Book not deleted: %v - isbn: %d\n", err, isbn)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
