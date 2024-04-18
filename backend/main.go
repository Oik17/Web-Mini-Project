package main

import (
	"github.com/Oik17/WebMiniProj/database"
	"github.com/gofiber/fiber/v2"
)
func main() {
	database.Connect()
	app := fiber.New()
	app.Listen(":3000")
}
