package main

import (
	"github.com/gofiber/fiber"
	"sopordiversao/controllers"
	"sopordiversao/middlewares"
)

func main() {
	app := fiber.New()

	// GET /john
	app.Get("/ping", controllers.Ping)
	//middleware
	app.Post("/login", controllers.Login)
	app.Use("/diversao", middlewares.Authentication)
	app.Static("/diversao", "html/index.html")
	app.Listen(3000)
}
