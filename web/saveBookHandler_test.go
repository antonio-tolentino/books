package web_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/models"
	"books/storage"
	"books/web"
)

var _ = Describe("SaveBookHandler", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	book := models.Book{Isbn: 123456,
		Title:  "Book Three updated",
		Author: &models.Author{Firstname: "Adam", Lastname: "Smith"}}

	// Test SaveBookHandler Success
	Describe("SaveBookHandler Success", func() {
		Context("When book is saved", func() {
			It("IT should return a book", func() {
				jsonData, _ := json.Marshal(book)
				req, _ := http.NewRequest("PUT", "/api/book", bytes.NewBuffer(jsonData))
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})
	})

	// Test SaveBookHandler Fail isbn invalid
	Describe("SaveBookHandler Fail", func() {
		Context("When book is not saved", func() {
			It("IT should return an error", func() {
				book2 := `{"isbn": "123456","title": "Book Three updated", "author": {"firstname": "Adam", "lastname": "Smith"}}`
				req, _ := http.NewRequest("PUT", "/api/book", bytes.NewBuffer([]byte(book2)))
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})

	// Test SaveBookHandler Fail fields empty
	Describe("SaveBookHandler Fail", func() {
		Context("When book is not saved", func() {
			It("IT should return an error", func() {
				book3 := `{"isbn": 123456,"title": "", "author": {"firstname": "", "lastname": ""}}`
				req, _ := http.NewRequest("PUT", "/api/book", bytes.NewBuffer([]byte(book3)))
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})
})
