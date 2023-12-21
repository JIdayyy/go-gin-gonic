package routes

import (
	"log"
	"net/http"

	. "local/database/users"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, GetUsers())
	})

	users.POST("/", func(c *gin.Context){
		body := User{}
		err := c.ShouldBind(&body)

		if(err != nil) {
			log.Fatal(err)
		}

		c.JSON(http.StatusCreated, CreateUser(body))
	})
}