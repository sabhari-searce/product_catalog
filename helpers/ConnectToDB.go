package helpers

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	connection_string := "user=service-pc-api host=localhost password=password sslmode=disable dbname=product-catalog"
	var err error
	db, err = sql.Open("postgres", connection_string)
	if err != nil {
		panic("Error in opening connection to database!!")
	}

}

func ConnectToDB() *sql.DB {

	return db
}
