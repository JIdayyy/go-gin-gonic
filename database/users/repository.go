package database

import (
	database "local/database"
	"log"
)

type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
}

func CreateUser(user User) User {
	var createdUSer User

	sqlStatement := `
      INSERT INTO users (username,email,password)
      VALUES ($1, $2, $3) RETURNING username, email, password;`

	value, err := database.PG.Prepare(sqlStatement)

	if err != nil {
		log.Fatal(err)
	}

	defer value.Close()

	value.QueryRow(user.Username, user.Email, user.Password).Scan(&createdUSer.Id, &createdUSer.Username, &createdUSer.Email, &createdUSer.Password, &createdUSer.Created_at)

	return user
}

func GetUsers() []User {
	rows, err := database.PG.Query("SELECT * FROM users;")

	users := make([]User, 0)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Created_at)

		if err != nil {
			log.Fatal(err)
		}

		newUser := User{
			Id:         user.Id,
			Username:   user.Username,
			Email:      user.Email,
			Password:   user.Password,
			Created_at: user.Created_at,
		}

		users = append(users, newUser)
	}

	return users
}

func DeleteUser(id string) string {
	result := ""
	sqlStatement := `DELETE FROM users WHERE id = $1`

	database.PG.Exec(sqlStatement, id)

	print("RESULT", result)

	return "OK"
}
