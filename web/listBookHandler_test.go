package web_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/storage"
	"books/web"
)

var _ = Describe("ListBookHandler", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	// Test ListBookHandler Success
	Describe("ListBookHandler Success", func() {
		Context("When book is found", func() {
			It("IT should return a book", func() {
				req, _ := http.NewRequest("GET", "/api/book/123456", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})

	// Test ListBookHandler Fail
	Describe("ListBookHandler Fail", func() {
		Context("When book is not found", func() {
			It("IT should return an error", func() {
				req, _ := http.NewRequest("GET", "/api/book/0", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

	// Test ListBookHandler Fail isbn conversion
	Describe("ListBookHandler Fail isbn conversion", func() {
		Context("When book is not found", func() {
			It("IT should return an error", func() {
				req, _ := http.NewRequest("GET", "/api/book/123abc", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
		})
	})

})
