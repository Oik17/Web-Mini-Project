package controllers

import (
	"context"
	"net/http"

	"github.com/Oik17/WebMiniProj/database"
	"github.com/Oik17/WebMiniProj/models"
	"github.com/gofiber/fiber/v2"
)

func CreateFoodItem(c *fiber.Ctx) error {
	db := database.DB.Db
	var menu models.Menu
	err := c.BodyParser(&menu)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Invalid input",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	query := `INSERT INTO menu (name, description, price) VALUES ($1, $2, $3)`
	_, err = db.ExecContext(
		context.Background(),
		query,
		menu.Name,
		menu.Description,
		menu.Price,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload menu",
			"data":    err.Error(),
			"status":  "false",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Menu Created",
		"data":    menu,
	})
}

func GetMenu(c *fiber.Ctx) error {
	db := database.DB.Db
	var menu []models.Menu
	query := `SELECT * FROM menu`

	err := db.Select(&menu, query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch menu",
			"data":    err.Error(),
			"status":  "false",
		})
	}
	return c.Status(http.StatusAccepted).JSON(menu)
}
