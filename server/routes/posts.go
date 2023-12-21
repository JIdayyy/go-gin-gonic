package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPostRoutes(rg *gin.RouterGroup) {
	posts := rg.Group("/posts")

	posts.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "POSTS")
	})
}