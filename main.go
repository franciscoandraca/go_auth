package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "usuaeio:Contraseña@/base_dattos?parseTime=false"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// Manejar el error
		panic("Error: no se pudo conectar a la base de datos")
	} else {
		// La conexión a la base de datos MySQL se ha establecido con éxito
		fmt.Println("La conexión a la base de datos MySQL se ha establecido con éxito")
		fmt.Println(db)
	}

	app := fiber.New()

	app.Get("/", home)

	app.Listen(":8000")
}

func home(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
