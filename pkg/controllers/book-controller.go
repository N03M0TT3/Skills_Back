package controllers

import "net/http"

// handler function that gets all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all books"))
}

// handler function that gets a single book
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get a single book"))
}

// handler function that creates a book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a book"))
}

// handler function that updates a book
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update a book"))
}

// handler function that deletes a book
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete a book"))
}
