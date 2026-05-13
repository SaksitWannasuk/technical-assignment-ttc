package main

import (
	"log"

	"technical-assignment/backend/internal/database"
	"technical-assignment/backend/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	app.Use(cors.New())

	routes.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Backend is running",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
