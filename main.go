package main

import (
	"st-portier-be/config"
	"st-portier-be/routes"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	r := routes.InitRoutes()
	r.Run(":8080")
}
