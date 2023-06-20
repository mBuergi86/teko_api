package Router

import (
	"awesomeProject/Handler"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	app.Get("/", Handler.Todos)
	app.Get("/:id", Handler.Todo)
	app.Post("/", Handler.CreateTodo)
	app.Put("/:id", Handler.UpdateTodo)
	app.Delete("/:id", Handler.DeleteTodo)
}
