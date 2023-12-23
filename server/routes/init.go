package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func InitRoutes() {
	v1 := router.Group("/v1")

	addUserRoutes(v1)

	router.Run(":4100")
}
