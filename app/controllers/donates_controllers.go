package controllers

import (
	"donation_app/app/models"
	"donation_app/app/queries"
	"donation_app/app/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateDonateHandler(c *fiber.Ctx) error {
	donate := models.Donate{}

	donate.CreatedAt = time.Now()
	donate.LoginWhoDonate = c.FormValue("loginwhodonate")
	donate.LoginToDonate = c.FormValue("logintodonate")
	donate.Message = c.FormValue("message")
	donate.Summary, _ = strconv.Atoi(c.FormValue("summary"))

	if err := queries.InsertDonate(donate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"donate": donate,
	})
}

func ListAllDonatesHandler(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	login := c.Params("login")

	donates, err := queries.ListAllDonates(login)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "books were not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(donates),
		"donates": donates,
	})
}
