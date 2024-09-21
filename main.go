package main

import (
	"donation_app/app/configs"
	"donation_app/app/middleware"
	"donation_app/app/routes"
	"donation_app/app/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.AdminRoutes(app)

	utils.StartServerWithGracefulShutdown(app)
}
