package database

import (
	"database/sql"
	database "local/database"
	"log"
)

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Email  string `json:"email"`
	Password string `json:"password"`
	Created_at  string `json:"created_at"`
}

func CreateUser(user User) sql.Result {
	sqlStatement := `
      INSERT INTO users (id,name)
      VALUES ($1, $2)`

	value,err := database.PG.Exec(sqlStatement,user.Id,user.Username)

	if(err != nil ){
		log.Fatal(err)
	}

    return value
}


func GetUsers() []User {
	rows, err := database.PG.Query("SELECT * FROM users;")

		users := make([]User,0)

		if(err != nil) {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			var user User
		    
			err := rows.Scan(&user.Id,&user.Username, &user.Email, &user.Password,&user.Created_at)

			if(err != nil) {
				log.Fatal(err)
			}

			newUser :=  User{
				Id: user.Id,
				Username: user.Username,
				Email: user.Email,
				Password: user.Password,
				Created_at: user.Created_at,
			}

		    users = append(users, newUser)
		}

		return users
}