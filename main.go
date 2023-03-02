package main

import (
	"task-5-vix-btpns-SuburFirmansyah/controllers"
	"task-5-vix-btpns-SuburFirmansyah/router"
)

func main() {
	server := controllers.Server{}
	server.Initialize()
	router.InitRoutes(&server)
	server.Run(8000)
}
