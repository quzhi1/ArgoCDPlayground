package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create fiber server
	app := fiber.New()

	// Define routes
	app.Get("/hello", hello)

	// Listen and serve
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

// getUser return user name by id
func hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"hello": "world"})
}
