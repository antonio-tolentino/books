package web_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/storage"
	"books/web"
)

var _ = Describe("DeleteBookHandler", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	// Test DeleteBookHandler Success
	Describe("DeleteBookHandler Success", func() {
		Context("When book is deleted", func() {
			It("IT should return a book", func() {
				req, _ := http.NewRequest("DELETE", "/api/book/123456", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})
	})

	// Test DeleteBookHandler Fail book not found
	Describe("DeleteBookHandler Fail", func() {
		Context("When book is not found", func() {
			It("IT should return an error", func() {
				req, _ := http.NewRequest("DELETE", "/api/book/0", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

	// Test DeleteBookHandler Fail isbn conversion
	Describe("DeleteBookHandler Fail isbn conversion", func() {
		Context("When book is not found", func() {
			It("IT should return an error", func() {
				req, _ := http.NewRequest("DELETE", "/api/book/123abc", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

})
