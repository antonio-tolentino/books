package storage_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/models"
	"books/storage"
)

var _ = Describe("BooksMemory", func() {

	// Init
	ctx := context.Background()
	booksRepo := storage.NewBooksMemoryStorage()

	// Test GetByIsbn Success
	Describe("GetByIsbn Success", func() {
		Context("When book is found", func() {
			It("IT should return a book", func() {
				book, err := booksRepo.GetByIsbn(ctx, 123456)
				Expect(err).To(BeNil())
				Expect(book.Isbn).To(Equal(123456))
			})
		})
	})

	// Test GetByIsbn Fail
	Describe("GetByIsbn Fail", func() {
		Context("When book is not found", func() {
			It("IT should return an error", func() {
				_, err := booksRepo.GetByIsbn(ctx, 0)
				Expect(err).To(Equal(storage.ErrBookNotFound))
			})
		})
	})

	// Test GetAll
	Describe("GetAll", func() {
		Context("When books are found", func() {
			It("IT should return a list of books", func() {
				books := booksRepo.GetAll(ctx)
				Expect(len(books)).To(Equal(4))
			})
		})
	})

	// Test Save success
	Describe("Save success", func() {
		Context("When book is saved", func() {
			It("IT should return a book", func() {
				book := models.Book{Isbn: 123456,
					Title:  "Book Three updated",
					Author: &models.Author{Firstname: "Adam", Lastname: "Smith"}}
				err := booksRepo.Save(ctx, book)
				Expect(err).To(BeNil())
			})
		})
	})

	// Test Save fail due to empty fields
	Describe("Save fail", func() {
		Context("When book is not saved", func() {
			It("IT should return an error", func() {
				book := models.Book{Isbn: 123456,
					Title:  "",
					Author: &models.Author{Firstname: "", Lastname: ""}}
				err := booksRepo.Save(ctx, book)
				Expect(err).To(Equal(storage.ErrAllFieldsRequired))
			})
		})
	})

	// Test Delete success
	Describe("Delete", func() {
		Context("When book is deleted", func() {
			It("IT should return a book", func() {
				err := booksRepo.Delete(ctx, 987654)
				Expect(err).To(BeNil())
			})
		})
	})

	// Test Delete fail
	Describe("Delete fail", func() {
		Context("When book is not found", func() {
			It("IT should return an error", func() {
				err := booksRepo.Delete(ctx, 0)
				Expect(err).To(Equal(storage.ErrBookNotFound))
			})
		})
	})

})
