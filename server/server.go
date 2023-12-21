package server

import (
	routes "github.com/JIdayyy/go-gin-gonic/server/routes"
)


func RunServer(){
	routes.InitRoutes()
}