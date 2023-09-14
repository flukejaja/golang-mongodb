package routes

import (
    "github.com/gofiber/fiber/v2"
    "myapp/controllers"
)

func CustomerRoutes(api fiber.Router) {
    api.Get("/user/:userId", controllers.GetAUser)
    api.Get("/users", controllers.GetAllUsers)

}