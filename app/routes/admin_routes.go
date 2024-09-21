package routes

import (
	"donation_app/app/controllers"
	"donation_app/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(a *fiber.App) {
	route := a.Group("/api/v1/admin")
	// ADMIN JWT-TOKEN UNLIMITED TIME -
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzYyNTc5NTJ9.aAkHcoBExc3UXnrrdlkNgIDK5TRDzewZOLbc2aCdJhM
	route.Get("/users", middleware.JWTProtected(), controllers.GetAllUsersHandler)

	route.Get("/donates", middleware.JWTProtected(), controllers.GetAllDonatesHandler)

	route.Delete("/delete/donate/:id", middleware.JWTProtected(), controllers.DeleteDonateHandler)
}
