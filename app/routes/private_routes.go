package routes

import (
	"donation_app/app/controllers"
	"donation_app/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	//route.Get("/users", middleware.JWTProtected(), controllers.GetAllUsersHandler)

	route.Get("/donates/:login", middleware.JWTProtected(), controllers.ListAllDonatesHandler)
}
