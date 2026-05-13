package handlers

import (
	"time"

	"technical-assignment/backend/internal/database"
	"technical-assignment/backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CreatePersonRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
	Address   string `json:"address"`
}

type PersonResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	BirthDate time.Time `json:"birthDate"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreatePerson(c *fiber.Ctx) error {
	var req CreatePersonRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid birth date format",
		})
	}

	person := models.Person{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: birthDate,
		Address:   req.Address,
	}

	result := database.DB.Create(&person)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create person",
		})
	}

	response := PersonResponse{
		ID:        person.ID,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: person.BirthDate,
		Age:       CalculateAge(person.BirthDate),
		Address:   person.Address,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	}

	return c.Status(201).JSON(response)
}

func GetPeople(c *fiber.Ctx) error {
	var people []models.Person
	result := database.DB.Order("id asc").Find(&people)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to get people",
		})
	}

	responses := []PersonResponse{}

	for _, person := range people {
		responses = append(responses, PersonResponse{
			ID:        person.ID,
			FirstName: person.FirstName,
			LastName:  person.LastName,
			BirthDate: person.BirthDate,
			Age:       CalculateAge(person.BirthDate),
			Address:   person.Address,
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
		})
	}

	return c.JSON(responses)
}

func GetPersonByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var person models.Person

	result := database.DB.First(&person, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Person not found",
		})
	}

	response := PersonResponse{
		ID:        person.ID,
		FirstName: person.FirstName,
		LastName:  person.LastName,
		BirthDate: person.BirthDate,
		Age:       CalculateAge(person.BirthDate),
		Address:   person.Address,
		CreatedAt: person.CreatedAt,
		UpdatedAt: person.UpdatedAt,
	}

	return c.JSON(response)
}

func CalculateAge(birthDate time.Time) int {
	now := time.Now()

	age := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}
