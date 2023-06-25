package main

import (
	"awesomeProject/middlewares"
	"awesomeProject/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	app.Get("/", monitor.New())
	middlewares.Middleware(app)
	router.Routing(app)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
