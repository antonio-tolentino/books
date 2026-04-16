package web_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/storage"
	"books/web"
)

var _ = Describe("HomeHandler", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	// Test HomeHandler Success
	Describe("HomeHandler Success", func() {
		Context("When / route is requested", func() {
			It("IT should return a book", func() {
				req, _ := http.NewRequest("GET", "/", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})

})
