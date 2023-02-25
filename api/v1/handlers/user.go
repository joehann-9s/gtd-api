package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"ok":      true,
		"message": "GTD API: Getting All user",
	})
}
