package main

import (
	"final/config"
	"final/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	r.Run(":8080")
}
