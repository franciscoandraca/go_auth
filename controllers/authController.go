package controllers

import (
	"go_auth/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var user models.User

	user.FirstName = "Francisco"
	user.LastName = "Andraca"
	user.Email = "email@gmail.com"
	user.Password = "secret"

	return c.JSON(user)
}
