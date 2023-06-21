package main

import (
	"awesomeProject/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Routing(app)

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
