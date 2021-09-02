package main

import (
	v1 "github.com/Fix-Pay/abc_brasil/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	crs := cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,DELETE,PUT,OPTIONS",
		AllowCredentials: true,
		AllowHeaders:     "*",
	})

	app.Use(crs)

	v1.Routes(app)

	app.Listen(":5000")
}
