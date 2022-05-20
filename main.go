package main

import (
	"crud/models"
	"crud/routes"
)

func main() {
	r := routes.SetupRoutes()
	models.SetUpDB()
	r.Run("localhost:8008")
}
