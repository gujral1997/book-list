package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/gujral1997/book-list/utils"
	"github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("SQL_URL"))

	utils.LogFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	utils.LogFatal(err)

	err = db.Ping()
	utils.LogFatal(err)

	log.Println("DB Started")
	return db
}
