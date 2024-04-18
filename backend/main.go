package main

import (
	"github.com/Oik17/WebMiniProj/controllers"
	"github.com/Oik17/WebMiniProj/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()

	app.Use(cors.New())
	app.Post("/create-food", controllers.CreateFoodItem)
	app.Get("/get-menu", controllers.GetMenu)

	app.Listen(":3000")
}
