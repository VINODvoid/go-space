package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (store *Store) Bookshandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(store.Books)
	case "POST":
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		created := store.Add(book.Title, book.Author, book.Year)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	}
}

func (store *Store) Bookhandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bookId := r.URL.Path[len("/books/"):]
	id, err := strconv.Atoi(bookId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintln(w, "Check the ID first !!")
		return
	}
	switch r.Method {
	case "GET":
		book, present := store.GetByID(id)
		if present {
			json.NewEncoder(w).Encode(book)
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintln(w, "book not found")
		}
	case "DELETE":
		deleted := store.Delete(id)
		if deleted {
			fmt.Fprintln(w, "Success !! deleted")
		} else {
			w.WriteHeader(http.StatusNotFound) // 404
		}

	}
}

func main() {
	store := &Store{}
	store.Add("The Go Programming Language", "Donovan", 2015)
	store.Add("Clean Code", "Martin", 2008)
	fmt.Println("Server is running on http://localhost:8080")
	fmt.Println("Available paths :\n1. /books\n2. /books/{id}\n3. DELETE /books/{id}")

	// Handlers for routes

	// 1.  Get details of all books
	http.HandleFunc("/books", store.Bookshandler)

	// 2. Get Book detailes by ID
	http.HandleFunc("/books/", store.Bookhandler)

	http.ListenAndServe(":8080", nil)
}
