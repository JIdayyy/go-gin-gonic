package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func createInstance() *sql.DB {
	connStr := "postgres://postgres:63b9c3f55ab687dd@5.39.75.106/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

    if err != nil {
		log.Fatal(err)
	}


	return db
}

var PG = createInstance()