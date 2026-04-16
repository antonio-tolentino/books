package web

import (
	"books/storage"

	"github.com/gorilla/mux"
)

type StorageHandler struct {
	Storage storage.BooksRepository
}

func NewStorageHandler(storage storage.BooksRepository) *StorageHandler {
	return &StorageHandler{
		Storage: storage,
	}
}

func (s *StorageHandler) RegisterRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/api/health", healthHandler).Name("helth-check").Methods("GET")
	r.HandleFunc("/api/books", s.listAllBooksHandler).Methods("GET")
	r.HandleFunc("/api/book/{isbn}", s.listBookHandler).Methods("GET")
	r.HandleFunc("/api/book", s.saveBookHandler).Methods("PUT")
	r.HandleFunc("/api/book/{isbn}", s.deleteBookHandler).Methods("DELETE")
	return r
}
