package main

import (
	"awesomeProject/Router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	Router.Routing(app)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
