package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

func Middleware(app *fiber.App) fiber.Handler {
	app.Use(logger.New())
	app.Use(recover2.New())

	return func(ctx *fiber.Ctx) error {

		fmt.Printf("INFO: %s %s\n", ctx.Method(), ctx.Path())

		return ctx.Next()
	}
}
