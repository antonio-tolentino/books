package web_test

import (
	"books/storage"
	"books/web"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BooksStorageHandler", func() {

	// Init
	booksRepo := storage.NewBooksMemoryStorage()
	s := web.NewStorageHandler(booksRepo)

	// Test RegisterRoutes Success
	Describe("RegisterRoutes Success", func() {
		// Registered routes
		r := s.RegisterRoutes()

		Context("When routes are registered", func() {
			It("IT should return a router", func() {
				Expect(r).NotTo(BeNil())
			})
		})

		// /api/health route was registered
		Describe("Health route", func() {
			Context("When /api/health route is registered", func() {
				It("IT should return a router", func() {
					path, _ := r.GetRoute("helth-check").GetPathTemplate()
					Expect(path).To(Equal("/api/health"))
				})
			})
		})
	})
})
