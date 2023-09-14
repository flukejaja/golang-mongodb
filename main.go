package main

import (
    "github.com/gofiber/fiber/v2"
    "myapp/routes"
	"myapp/config"
)

func main() {
    app := fiber.New()

	configs.ConnectDB()

	api := app.Group("/api")

	routes.CustomerRoutes(api)

    app.Listen(":3000")
}

