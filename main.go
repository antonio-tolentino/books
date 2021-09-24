package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Status struct
type Status struct {
	Status string `json:"status"`
}

// Status struct
type OutBoundAddress struct {
	IP string `json:"ip"`
}

//Init books var as a slice Book Struct
var books []Book

// ContentType header
const ContentType = "Content-Type"

// ApplicationJSON header
const ApplicationJSON = "application/json"

// APIBooksID path
const APIBooksID = "/api/books/{id}"

//APIBooks path
const APIBooks = "/api/books"

//APIHealth path
const APIHealth = "/api/health"

//OutBoundIP path
const OutBoundIP = "/api/outboundip"

// Custom response header
const myResponse = "myResponse"
const myResponseMsg = "Hello"

// Show OutBoundIP
func outBoundIP(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://ifconfig.me/ip")
	if err != nil {
		log.Fatalln("Error making GET: ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading body: ", err)
	}

	log.Printf("Output ip = %s", string(body))

	w.Header().Set(ContentType, ApplicationJSON)
	w.Header().Add(myResponse, myResponseMsg)
	var address OutBoundAddress
	address.IP = string(body)
	json.NewEncoder(w).Encode(address)
}

// Show API Health
func health(w http.ResponseWriter, r *http.Request) {

	id := uuid.New()
	log.Printf("%s - method   = %s", id, r.Method)
	log.Printf("%s - url      = %s", id, r.URL)
	log.Printf("%s - protocol = %s", id, r.Proto)
	h, err := json.MarshalIndent(r.Header, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - header   = \n%s\n", id, string(h))

	w.Header().Set(ContentType, ApplicationJSON)
	w.Header().Add(myResponse, myResponseMsg)
	var status Status
	status.Status = "I am fine!"
	json.NewEncoder(w).Encode(status)
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {

	id := uuid.New()
	log.Printf("%s - method   = %s", id, r.Method)
	log.Printf("%s - url      = %s", id, r.URL)
	log.Printf("%s - protocol = %s", id, r.Proto)
	h, err := json.MarshalIndent(r.Header, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - header   = \n%s\n", id, string(h))

	w.Header().Set(ContentType, ApplicationJSON)
	w.Header().Add(myResponse, myResponseMsg)
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {

	id := uuid.New()
	log.Printf("%s - method    = %s", id, r.Method)
	log.Printf("%s - url       = %s", id, r.URL)
	log.Printf("%s - protocol  = %s", id, r.Proto)
	p, err := json.MarshalIndent(mux.Vars(r), "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - parameters = \n%s\n", id, string(p))

	h, err := json.MarshalIndent(r.Header, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - header     = \n%s\n", id, string(h))

	w.Header().Set(ContentType, ApplicationJSON)
	w.Header().Add(myResponse, myResponseMsg)

	params := mux.Vars(r) // Get params

	// loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create new book
func createBook(w http.ResponseWriter, r *http.Request) {

	id := uuid.New()
	log.Printf("%s - method    = %s", id, r.Method)
	log.Printf("%s - url       = %s", id, r.URL)
	log.Printf("%s - protocol  = %s", id, r.Proto)
	h, err := json.MarshalIndent(r.Header, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - header     = \n%s\n", id, string(h))

	w.Header().Set(ContentType, ApplicationJSON)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	books = append(books, book)

	w.Header().Add(myResponse, myResponseMsg)

	json.NewEncoder(w).Encode(book)

	b, err := json.MarshalIndent(book, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s - body       = \n%s\n", id, string(b))
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request) {

	id := uuid.New()
	log.Printf("%s - method     = %s", id, r.Method)
	log.Printf("%s - url        = %s", id, r.URL)
	log.Printf("%s - protocol   = %s", id, r.Proto)
	p, err := json.MarshalIndent(mux.Vars(r), "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s - parameters = \n%s\n", id, string(p))
	h, err := json.MarshalIndent(r.Header, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s - headers    = \n%s\n", id, string(h))

	w.Header().Set(ContentType, ApplicationJSON)
	w.Header().Add(myResponse, myResponseMsg)

	params := mux.Vars(r) // Get params
	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)

			b, err := json.MarshalIndent(book, "", "\t")
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%s - body       = \n%s\n", id, string(b))

			return
		}
	}
	json.NewEncoder(w).Encode(books)

}

// Delete Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	id := uuid.New()
	log.Printf("%s - method     = %s", id, r.Method)
	log.Printf("%s - url        = %s", id, r.URL)
	log.Printf("%s - protocol   = %s", id, r.Proto)
	p, err := json.MarshalIndent(mux.Vars(r), "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s - parameters = \n%s\n", id, string(p))
	h, err := json.MarshalIndent(r.Header, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s - header     = \n%s\n", id, string(h))

	w.Header().Set(ContentType, "application/json")
	w.Header().Add(myResponse, myResponseMsg)

	params := mux.Vars(r) // Get params
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func mockData() {
	// Mock Data - @todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "448743",
		Title:  "Book One",
		Author: &Author{Firstname: "Agatha", Lastname: "Black"}})
	books = append(books, Book{ID: "2", Isbn: "556798",
		Title:  "Book Two",
		Author: &Author{Firstname: "Steve", Lastname: "White"}})
}

func createRouter() (router *mux.Router) {

	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc(APIHealth, health).Methods("GET")
	r.HandleFunc(OutBoundIP, outBoundIP).Methods("GET")
	r.HandleFunc(APIBooks, getBooks).Methods("GET")
	r.HandleFunc(APIBooksID, getBook).Methods("GET")
	r.HandleFunc(APIBooks, createBook).Methods("POST")
	r.HandleFunc(APIBooksID, updateBook).Methods("PUT")
	r.HandleFunc(APIBooksID, deleteBook).Methods("DELETE")

	return r
}

func main() {
	// create mock data
	mockData()

	// Start server
	log.Println("Change code")
	log.Println("Starting server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", createRouter()))
}
