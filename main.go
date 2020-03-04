package main

// install the package mux and import it
// https://github.com/gorilla/mux

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Writer `json:"author"`
}

type Writer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// Init books var as a slice Book struct
var books []Book

func main() {
	// Init the http router
	router := mux.NewRouter()

	// Test endpoint
	router.HandleFunc("/test", testAPI).Methods("GET")

	books = append(books, Book{ID: "1", Title: "Madame Bovary", Author: &Writer{FirstName: "Gustave", LastName: "Flaubert"}})
	books = append(books, Book{ID: "2", Title: "Les miserables", Author: &Writer{FirstName: "Victor", LastName: "Hugo"}})

	// Route handlers & endpoints
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	// router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// Start the server
	// http.ListenAndServe(":5000", router)

	// To be able to log an error if it occurs
	log.Fatal(http.ListenAndServe(":5000", router))

}

func testAPI(w http.ResponseWriter, r *http.Request) {
	// w.Write("all right")
	json.NewEncoder(w).Encode("all right")
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	fmt.Println("params: ", params)
	// Loop through books and find one with the id from the params
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	// json.NewEncoder(w).Encode(&Book{})
}

// Add new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000)) // Mock ID - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
