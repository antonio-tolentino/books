package web_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/storage"
	"books/web"
)

var _ = Describe("GetAll", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	// Test GetAll Success
	Describe("GetAll Success", func() {
		Context("When books are found", func() {
			It("IT should return a list of books", func() {
				req, _ := http.NewRequest("GET", "/api/books", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
