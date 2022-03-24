package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendStatus(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"Message": "Hello World!",
		})
	})

	app.Listen(":3000")
}
