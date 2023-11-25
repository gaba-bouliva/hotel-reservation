package api

import (
	"github.com/gaba-bouliva/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Jane",
		LastName: "Doe",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Jane Doe")
}