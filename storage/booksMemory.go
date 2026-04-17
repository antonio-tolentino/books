package storage

import (
	"books/models"
	"context"
	"errors"
	"log/slog"
	"sync"
)

// custom book errors
var (
	ErrBookNotFound      = errors.New("Book isbn not found")
	ErrAllFieldsRequired = errors.New("All fields are required")
)

// BooksMemoryStorage
type BooksMemoryStorage struct {
	mu   sync.RWMutex // Protects the map from concurrent access
	data map[int]models.Book
}

// NewBooksMemoryStorage returns a new in-memory storage
func NewBooksMemoryStorage() *BooksMemoryStorage {

	// make a new map and call mockData
	books := mockData()
	slog.Info("Books memory storage initialized...")

	return &BooksMemoryStorage{
		data: books,
	}
}

// implement GetByIsbn
func (bs *BooksMemoryStorage) GetByIsbn(ctx context.Context, isbn int) (models.Book, error) {
	bs.mu.RLock()
	defer bs.mu.RUnlock()

	book, ok := bs.data[isbn]
	if !ok {
		return models.Book{}, ErrBookNotFound
	}
	return book, nil
}

// implement GetAll
func (bs *BooksMemoryStorage) GetAll(ctx context.Context) []models.Book {
	bs.mu.RLock()
	defer bs.mu.RUnlock()
	books := make([]models.Book, 0, len(bs.data))
	for _, v := range bs.data {
		books = append(books, v)
	}

	return books
}

// implement Save
func (bs *BooksMemoryStorage) Save(ctx context.Context, book models.Book) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	if book.Title == "" || book.Author.Firstname == "" || book.Author.Lastname == "" {
		return ErrAllFieldsRequired
	}
	bs.data[book.Isbn] = book
	return nil
}

// implement Delete
func (bs *BooksMemoryStorage) Delete(ctx context.Context, isbn int) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	// check if book exists
	_, ok := bs.data[isbn]
	if ok {
		delete(bs.data, isbn)
		return nil
	}

	return ErrBookNotFound
}
