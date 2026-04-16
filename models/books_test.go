package models_test

import (
	"books/models"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Books", func() {

	var b *models.Book

	BeforeEach(func() {
		b = &models.Book{Isbn: 123456,
			Title:  "Book One",
			Author: &models.Author{Firstname: "Agatha", Lastname: "Black"}}
	})

	Describe("Isbn Validation", func() {
		Context("When Isbn is valid", func() {
			It("IT should return true", func() {
				Expect(b.Isbn).To(BeNumerically(">", 100000))
			})
		})

		Context("When Isbn is not valid", func() {
			It("IT should return false", func() {
				b.Isbn = 0
				Expect(b.Isbn).To(BeNumerically("<", 100000))
			})
		})
	})

	Describe("String", func() {
		Context("When String is called", func() {
			It("IT should return a string", func() {
				Expect(b.String()).To(Equal("Book One by Agatha"))
			})
		})
	})

})
