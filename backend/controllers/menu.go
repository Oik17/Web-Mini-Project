package controllers

import (
	"net/http"

	"github.com/Oik17/WebMiniProj/database"
	"github.com/Oik17/WebMiniProj/models"
	"github.com/gofiber/fiber/v2"
)

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
