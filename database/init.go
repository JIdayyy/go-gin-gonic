package database

import (
	"database/sql"
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