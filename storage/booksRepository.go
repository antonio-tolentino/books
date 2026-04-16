package storage

import (
	"books/models"
	"context"
)

type BooksRepository interface {
	GetAll(ctx context.Context) []models.Book
	GetByIsbn(ctx context.Context, isbn int) (models.Book, error)
	Save(ctx context.Context, book models.Book) error
	Delete(ctx context.Context, isbn int) error
}
