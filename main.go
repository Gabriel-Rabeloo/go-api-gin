package main

import (
	"github.com/Gabriel-Rabeloo/go-api-gin/database"
	"github.com/Gabriel-Rabeloo/go-api-gin/routes"
)

func main() {
	database.Connect()

	routes.HandleRequests()
}
