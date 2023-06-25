package router

import (
	"awesomeProject/handler"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	todos := app.Group("/todos")
	//v1 := todos.Group("/v1")

	todos.Get("/", handler.Todos)
	todos.Get("/:id", handler.TodoID)
	todos.Post("/", handler.CreateTodo)
	todos.Put("/:id", handler.UpdateTodo)
	todos.Delete("/:id", handler.DeleteTodo)
}
