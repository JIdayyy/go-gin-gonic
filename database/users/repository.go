package database

import (
	"fmt"
	database "local/database"
	"log"
)

type User struct {
	Id         string `json:"id", validate:"required`
	Username   string `json:"username, validate:"required"`
	Email      string `json:"email, validate:"required"`
	Password   string `json:"password, validate:"required"`
	Created_at string `json:"created_at, validate:"required"`
}

// UserBodyPatch represents the body for partial updates
type UserBodyPatch struct {
	Username *string `json:"username,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func CreateUser(user User) User {
	var createdUSer User
	fmt.Println(("CREATE USER"), user)
	sqlStatement := `
      INSERT INTO users (username,email,password)
      VALUES ($1, $2, $3) RETURNING username, email;`

	value, err := database.PG.Prepare(sqlStatement)

	if err != nil {
		fmt.Println("ERROR", err)
		log.Fatal(err)
	}

	defer value.Close()

	value.QueryRow(user.Username, user.Email, user.Password).Scan(&createdUSer.Id, &createdUSer.Username, &createdUSer.Email, &createdUSer.Created_at)
	fmt.Println("CREATED USER", createdUSer)
	fmt.Println("CREATED USER", value)
	return user
}

func GetUsers() []User {
	rows, err := database.PG.Query("SELECT id, email, username, created_at FROM users;")

	users := make([]User, 0)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Email, &user.Username, &user.Created_at)

		if err != nil {
			log.Fatal(err)
		}

		newUser := User{
			Id:         user.Id,
			Username:   user.Username,
			Email:      user.Email,
			Created_at: user.Created_at,
		}

		users = append(users, newUser)
	}

	return users
}

func GetUser(id string) User {
	sqlStatement := `SELECT id, email, username, created_at FROM users WHERE id = $1;`

	var user User
	err := database.PG.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Email, &user.Username, &user.Created_at)

	if err != nil {
		fmt.Println("ERROR", err)
		log.Fatal(err)
	}

	fmt.Println("USER", user)
	return user
}

func DeleteUser(id string) string {
	sqlStatement := `DELETE FROM users WHERE id = $1`

	database.PG.Exec(sqlStatement, id)

	return "OK"
}

// partial update, only update the fields that are passed in the body
func PatchUser(id string, body UserBodyPatch) User {
	fmt.Println("BODY", body)

	sqlStatement := `UPDATE users SET username = COALESCE($1, username), email = COALESCE($2, email), password = COALESCE($3, password) WHERE id = $4 RETURNING id, username, email, password;`

	var user User
	err := database.PG.QueryRow(sqlStatement, body.Username, body.Email, body.Password, id).Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		fmt.Println("ERROR", err)
		log.Fatal(err)
	}

	fmt.Println("Updated User:", user)
	return user
}
