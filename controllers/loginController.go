package controllers

import (
	"github.com/gofiber/fiber"
	"net/http"
	"sopordiversao/authc"
)

func Login(c *fiber.Ctx) {
	if c.Query("user") == "admin" && c.Query("senha") == "123" {
		token, err := authc.GenerateJWT(86400)
		if err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Deu Ruim",
			})
			return
		}
		c.Status(http.StatusOK).JSON(fiber.Map{
			"token": token,
		})
		return
	}
}
