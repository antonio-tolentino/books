package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	//Mock data
	mockData()

	req, err := http.NewRequest("GET", "/api/books", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var allBooks []Book
	_ = json.NewDecoder(response.Body).Decode(&allBooks)

	if allBooks[0].Isbn != books[0].Isbn {
		t.Errorf("Test Failed: Isbn = %s, expected = %s", allBooks[0].Isbn, books[0].Isbn)
	}

}

func TestGetBook(t *testing.T) {

	//Mock data
	mockData()

	req, err := http.NewRequest("GET", "/api/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetBookNotFound(t *testing.T) {

	//Mock data
	mockData()

	req, err := http.NewRequest("GET", "/api/books/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var book Book
	_ = json.NewDecoder(response.Body).Decode(&book)

	if book.Isbn != "" {
		t.Errorf("Test Failed: Isbn = %s, expected = \"\"", book.Isbn)
	}

}

func TestCreateBook(t *testing.T) {

	var jsonStr = []byte(`{"id":"1","isbn":"457890","title":"New Book","author":{"firstname":"Neil","lastname":"Gray"}}`)

	req, err := http.NewRequest("POST", "/api/books", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeleteBook(t *testing.T) {

	books = append(books, Book{ID: "1", Isbn: "448743",
		Title:  "Book One",
		Author: &Author{Firstname: "Agatha", Lastname: "Black"}})

	req, err := http.NewRequest("DELETE", "/api/books/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateBook(t *testing.T) {

	//Mock data
	mockData()

	var jsonStr = []byte(`{"id":"1","isbn":"448743","title":"Updated Book","author":{"firstname":"Agatha","lastname":"Black"}}`)

	req, err := http.NewRequest("PUT", "/api/books/1", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateNotFound(t *testing.T) {

	//Mock data
	mockData()

	var jsonStr = []byte(`{"id":"5","isbn":"448743","title":"Updated Book","author":{"firstname":"Agatha","lastname":"Black"}}`)

	req, err := http.NewRequest("PUT", "/api/books/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func Test_health(t *testing.T) {
	//Mock data
	mockData()

	req, err := http.NewRequest("GET", "/api/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	router := createRouter()

	router.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
