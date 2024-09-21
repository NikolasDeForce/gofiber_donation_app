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

	login, err := queries.FindUserLogin(donate.LoginToDonate)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "user not found in db",
		})
	}

	if login.Login != donate.LoginToDonate {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "user is doesn't match to login",
		})
	}

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
			"error":   true,
			"msg":     "donates were not found",
			"count":   0,
			"donates": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(donates),
		"donates": donates,
	})
}

func GetAllDonatesHandler(c *fiber.Ctx) error {
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

	donates, err := queries.GetAllDonates()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "donates were not found",
			"count":   0,
			"donates": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(donates),
		"donates": donates,
	})
}

func DeleteDonateHandler(c *fiber.Ctx) error {
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

	id, _ := strconv.Atoi(c.Params("id"))

	ok := queries.DeleteDonate(id)
	if ok != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "donates were not found",
			"count":   0,
			"donates": nil,
		})
	}
	donates, err := queries.GetAllDonates()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "donates were not found",
			"count":   0,
			"donates": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(donates),
		"donate": donates,
	})
}
