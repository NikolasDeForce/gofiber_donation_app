package routes

import (
	"donation_app/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/register", controllers.CreateUserHandler)
	route.Post("/donate", controllers.CreateDonateHandler)
}
