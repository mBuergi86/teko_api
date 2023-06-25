package router

import (
	"awesomeProject/handler"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App) {
	todos := app.Group("/todos")
	//v1 := todos.Group("/v1", func(c *fiber.Ctx) error {
	//		c.Set("Version", "v1")
	//		return c.Next()
	//	})

	todos.Get("/", handler.Todos)
	todos.Get("/:id<regex(^[0-9]+$)}>", handler.TodoID)
	todos.Post("/", handler.CreateTodo)
	todos.Put("/:id<regex(^[0-9]+$)}>", handler.UpdateTodo)
	todos.Delete("/:id<regex(^[0-9]+$)}>", handler.DeleteTodo)
}
