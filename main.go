package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gujral1997/book-list/controllers"
	"github.com/gujral1997/book-list/driver"
	"github.com/subosito/gotenv"
)

// Init books var as a slice Book struct

var db *sql.DB

func init() {
	gotenv.Load()
}

func main() {
	db = driver.ConnectDB()
	// init router
	router := mux.NewRouter()
	controller := controllers.Controller{}
	// Route handlers
	router.HandleFunc("/api/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/api/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/api/books", controller.CreateBook(db)).Methods("POST")
	router.HandleFunc("/api/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controller.DeleteBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
