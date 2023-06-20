package Handler

import (
	"github.com/gofiber/fiber/v2"
)

func Todos(c *fiber.Ctx) error {
	return c.SendString("Hello Todos")
}

func Todo(c *fiber.Ctx) error {
	return c.SendString("Hello Todo")
}

func CreateTodo(c *fiber.Ctx) error {
	return c.SendString("Hello Create")
}

func UpdateTodo(c *fiber.Ctx) error {
	return c.SendString("Hello Update")
}

func DeleteTodo(c *fiber.Ctx) error {
	return c.SendString("Hello Delete")
}
