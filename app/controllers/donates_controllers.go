package controllers

import (
	"donation_app/app/models"
	"donation_app/app/queries"
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
