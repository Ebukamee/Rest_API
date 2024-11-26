package main

import (
	"encoding/json"
	"log"
	// "fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}
type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lasttname"`
}

// Init Book slice
var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Loop through
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	//Generate random number from 1 t0 10 million , convert it to string an set as Book id
	book.ID = strconv.Itoa(rand.Intn(10000000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// Init Mux Browser
	r := mux.NewRouter()
	books = append(books, Book{ID: "1", Isbn: "23456789", Title: "Great Wars", Author: &Author{FirstName: "Emeka", LastName: "Okorie"}})
	// EndPoints / Route Handlers

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", DeleteBook).Methods("DELETE")
	// RUN SERVER

	log.Fatal(http.ListenAndServe(":8000", r))
}
