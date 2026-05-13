package routes

import (
	"technical-assignment/backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/people", handlers.GetPeople)
	api.Get("/person/:id", handlers.GetPersonByID)
	api.Post("/person", handlers.CreatePerson)
}
