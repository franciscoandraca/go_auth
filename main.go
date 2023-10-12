package main

import (
	"go_auth/database"
	"go_auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Conectar()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}
