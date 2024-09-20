package controllers

import (
	"donation_app/app/models"
	"donation_app/app/queries"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllUsersHandler(c *fiber.Ctx) error {
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
		"books": users,
	})
}

// func GetUserByEmailHandler(c *fiber.Ctx) error {
// 	email, err := uuid.Parse(c.Params("email"))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	db, err := db.DBConnection()
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   err.Error(),
// 		})
// 	}

// 	user, err := db.GetUserByEmail(email.String())
// 	if err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"error": true,
// 			"msg":   "user with the given email is not found",
// 			"user":  nil,
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"error": false,
// 		"msg":   nil,
// 		"user":  user,
// 	})
// }

func CreateUserHandler(c *fiber.Ctx) error {
	// now := time.Now().Unix()

	// claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// expires := claims.Expires

	// if now > expires {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	user := models.User{}

	user.CreatedAt = time.Now()
	user.Login = c.FormValue("login")
	user.Password = c.FormValue("password")
	user.Email = c.FormValue("email")

	// if err := c.BodyParser(user); err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	if err := queries.InsertUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
