package web_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"books/storage"
	"books/web"
)

var _ = Describe("HealthCheck", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)
	r := s.RegisterRoutes()

	// Test HealthCheck Status = UP
	Describe("HealthCheck Status = UP", func() {
		Context("When /api/health route is working", func() {
			It("IT should return status = UP", func() {

				req, _ := http.NewRequest("GET", "/api/health", nil)
				w := httptest.NewRecorder()
				r.ServeHTTP(w, req)
				h := web.Health{}
				_ = json.NewDecoder(w.Body).Decode(&h)
				Expect(h.Status).To(Equal("UP"))
			})
		})
	})

})
