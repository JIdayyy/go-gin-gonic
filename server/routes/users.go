package routes

import (
	"fmt"
	"log"
	"net/http"

	. "local/database/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": false, "message": err.Error()})
		return true // signal that there was an error and the caller should return
	}
	return false // no error, can continue
}

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, GetUsers())
	})

	users.GET("/:id", func(c *gin.Context) {
		id, _ := c.Params.Get("id")

		c.JSON(http.StatusOK, GetUser(id))
	})

	users.POST("/", func(c *gin.Context) {
		validator := validator.New()

		err := validator.Struct(User{})

		if err != nil {
			fmt.Println("ERROR USERS POST", err)
			log.Fatal(err)
		}

		body := User{}
		err2 := c.ShouldBind(&body)

		if err != nil {
			fmt.Println("ERROR USERS POST", err2)
			log.Fatal(err2)
		}

		c.JSON(http.StatusCreated, CreateUser(body))
	})

	users.DELETE("/:id", func(c *gin.Context) {
		id, err := c.Params.Get("id")

		if !err {
			log.Fatal(err)
		}

		DeleteUser(id)

	})

	users.PATCH("/:id", func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		fmt.Println("PATH USER REQUEST RECIEVED", id)
		body := UserBodyPatch{}

		err := c.ShouldBind(&body)

		if Error(c, err) {
			return // exit
		}

		newUser := PatchUser(id, body)

		c.JSON(http.StatusOK, newUser)

	})

	users.DELETE("/", func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		DeleteUser(id)

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
}
