package main

import (
	"encoding/json"
	"fmt"
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
var book []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
func getBook(w http.ResponseWriter, r *http.Request) {

}
func createBook(w http.ResponseWriter, r *http.Request) {

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {

}
func DeleteBook(w http.ResponseWriter, r *http.Request) {

}
func main() {
	// Init Mux Browser
	r := mux.NewRouter()
	book = append(book, Book{ID: "1", Isbn: "23456789", Title: "Great Wars", Author: &Author{FirstName: "Emeka", LastName: "Okorie"}})
	// EndPoints / Route Handlers

	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", UpdateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", DeleteBook).Methods("DELETE")
	// RUN SERVER
	log.Fatal(http.ListenAndServe(":8000", r))
}
