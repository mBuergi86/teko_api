package router

import (
	"awesomeProject/handler"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	app.Get("/todos", handler.Todos)
	app.Get("/todos/:id", handler.TodoID)
	app.Post("/todos", handler.CreateTodo)
	app.Put("/todos/:id", handler.UpdateTodo)
	app.Delete("/todos/:id", handler.DeleteTodo)
}
