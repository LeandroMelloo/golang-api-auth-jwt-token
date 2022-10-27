package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/leandromello/api-auth-jwt-token/databases"
	"github.com/leandromello/api-auth-jwt-token/routes"
)

func main() {
	databases.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":3000")
}