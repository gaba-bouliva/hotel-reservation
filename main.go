package main

import (
	"flag"

	api "github.com/gaba-bouliva/hotel-reservation/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", ":5000", "The API server listening address ")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo )
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)


}
func handleUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Jane doe"})
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "hello gopher!"})
}