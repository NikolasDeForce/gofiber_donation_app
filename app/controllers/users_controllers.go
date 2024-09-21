package controllers

import (
	"donation_app/app/models"
	"donation_app/app/queries"
	"donation_app/app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
	var adminToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzYyNTc5NTJ9.aAkHcoBExc3UXnrrdlkNgIDK5TRDzewZOLbc2aCdJhM"

	token, err := utils.VerifyToken(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if token.Raw != adminToken {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "given token not admin-token",
		})
	}

	users, err := queries.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "users were not found",
			"count": 0,
			"users": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(users),
		"users": users,
	})
}

func CreateUserHandler(c *fiber.Ctx) error {
	user := models.User{}

	user.CreatedAt = time.Now()
	user.Login = c.FormValue("login")
	user.Password = c.FormValue("password")
	user.Email = c.FormValue("email")

	if err := queries.InsertUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	token, err := utils.GenerateNewAccessToken()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"user":         user,
		"access_token": token,
	})
}
