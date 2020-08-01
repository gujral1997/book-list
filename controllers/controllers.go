package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gujral1997/book-list/models"
	bookrepository "github.com/gujral1997/book-list/repository/book"
)

type Controller struct{}

var books []models.Book

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		bookRepo := bookrepository.BookRepository{}
		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)

		bookRepo := bookrepository.BookRepository{}
		book = bookRepo.GetBook(db, book, params["id"])
		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) CreateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookrepository.BookRepository{}
		bookID := bookRepo.CreateBook(db, book)
		json.NewEncoder(w).Encode(bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookrepository.BookRepository{}
		rowsUpdated := bookRepo.UpdateBooks(db, book)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) DeleteBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		bookRepo := bookrepository.BookRepository{}
		rowDeleted := bookRepo.DeleteBooks(db, params["id"])
		json.NewEncoder(w).Encode(rowDeleted)
	}

}
