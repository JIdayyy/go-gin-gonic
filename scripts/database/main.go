package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func createInstance() *sql.DB {
	connStr := goDotEnvVariable("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	return db
}

var PG = createInstance()

func createTables(DB *sql.DB) {
	c, ioErr := os.ReadFile("/Users/julienabbadie/Dev/go-http/src/scripts/database/tables/tables.sql")

	if ioErr != nil {
		fmt.Print(ioErr)

	}

	sql := string(c)

	_, err := PG.Exec(sql)

	if err != nil {
		fmt.Print(err)
		log.Fatal("ERROR DURING TABLE CREATION")
	}
}

func populate(DB *sql.DB) {
	c, ioErr := os.ReadFile("/Users/julienabbadie/Dev/go-http/src/scripts/database/data/seeds.sql")

	if ioErr != nil {
		fmt.Print(ioErr)
	}

	sql := string(c)

	_, err := PG.Exec(sql)

	if err != nil {
		fmt.Print(err)
		log.Fatal("ERROR DURING DB POPULATE 🚀")
	}

}

func main() {
	createTables(PG)
	populate(PG)
}
