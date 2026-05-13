package main

import (
	"log"

	"technical-assignment/backend/internal/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Backend is running",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
