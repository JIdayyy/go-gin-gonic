package database

import (
	"database/sql"
	"log"

	database "github.com/JIdayyy/go-gin-gonic/database"
)

type User struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func CreateUser(user User) sql.Result {
	sqlStatement := `
      INSERT INTO test (id,name)
      VALUES ($1, $2)`

	value,err := database.PG.Exec(sqlStatement,user.Id,user.Name)

	if(err != nil ){
		log.Fatal(err)
	}

    return value
}


func GetUsers() []User {
	rows, err := database.PG.Query("SELECT * FROM test;")

		users := make([]User,0)

		if(err != nil) {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			var user User
		    
			err := rows.Scan(&user.Id,&user.Name)

			if(err != nil) {
				log.Fatal(err)
			}

			newUser :=  User{
				Id: user.Id,
				Name:user.Name,
			}

		    users = append(users, newUser)
		}

		return users
}