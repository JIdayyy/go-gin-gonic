package controllers

import (
	"net/http"

	usersRepository "github.com/JIdayyy/go-gin-gonic/database/users"
	"github.com/gin-gonic/gin"
)


func GetUsers(ctx *gin.Context)  {
	 ctx.JSON(http.StatusOK, usersRepository.GetUsers())
}

