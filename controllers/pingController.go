package controllers

import "github.com/gofiber/fiber"

func Ping(c *fiber.Ctx) {
	c.Send("Pong")
	return
}
