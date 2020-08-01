package bookrepository

import (
	"database/sql"

	"github.com/gujral1997/book-list/models"
	"github.com/gujral1997/book-list/utils"
)

type BookRepository struct{}

func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	utils.LogFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		utils.LogFatal(err)

		books = append(books, book)
	}
	return books
}

func (b BookRepository) GetBook(db *sql.DB, book models.Book, id string) models.Book {
	rows := db.QueryRow("select * from books where id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	utils.LogFatal(err)
	return book
}

func (b BookRepository) CreateBook(db *sql.DB, book models.Book) string {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&book.ID)
	utils.LogFatal(err)
	return book.ID
}

func (b BookRepository) UpdateBooks(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)
	utils.LogFatal(err)

	rowsUpdated, err := result.RowsAffected()
	utils.LogFatal(err)
	return rowsUpdated
}

func (b BookRepository) DeleteBooks(db *sql.DB, id string) int64 {
	result, err := db.Exec("delete from books where id=$1", id)
	utils.LogFatal(err)

	rowDeleted, err := result.RowsAffected()
	utils.LogFatal(err)
	return rowDeleted
}
