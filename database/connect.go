package database

import (
	"fmt"
	"go_auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conectar() {

	dsn := "usuaeio:Contraseña@/base_dattos?parseTime=false"

	conexion, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		// Manejar el error
		panic("Error: no se pudo conectar a la base de datos")
	} else {
		// La conexión a la base de datos MySQL se ha establecido con éxito
		fmt.Println("La conexión a la base de datos MySQL se ha establecido con éxito")
	}

	conexion.AutoMigrate(&models.User{})

}
