package server

import (
	routes "local/server/routes"
)


func RunServer(){
	routes.InitRoutes()
}