package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(app *fiber.App) {
	app.Use(logger.New())
	app.Use(recover2.New())
}
