package handler

import (
	"awesomeProject/entity"
	"awesomeProject/errors"
	"awesomeProject/persistence"
	"awesomeProject/utility"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Todos(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, err := persistence.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
	}
	return c.JSON(data.Todos())
}

func TodoID(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id := c.Params("id")
	if !utility.IsNumeric(id) {
		return c.Status(fiber.StatusBadRequest).JSON(errors.ErrInvalidID)
	}

	data, err := persistence.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
	}

	todo, err := data.TodoID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
	}

	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, err := persistence.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
	}

	var todo entity.Todo

	if err := json.Unmarshal(c.Body(), &todo); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(errors.ErrFileNotCreated)
	}

	return c.JSON(data.CreateTodo(todo))
}

func UpdateTodo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id := c.Params("id")
	if !utility.IsNumeric(id) {
		return c.Status(fiber.StatusBadRequest).JSON(errors.ErrTodoNotFound)
	}

	data, err := persistence.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
	}

	var todo entity.Todo

	if err := json.Unmarshal(c.Body(), &todo); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(errors.ErrNotImplemented)
	}

	newTodo, err := data.UpdateTodo(id, todo)

	if err != nil {
		if err.Error() == "todo not found" {
			return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
		}
		//file can not created
		return c.Status(fiber.StatusNotImplemented).JSON(errors.ErrFileNotCreated)
	}

	return c.JSON(newTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id := c.Params("id")
	if !utility.IsNumeric(id) {
		return c.Status(fiber.StatusBadRequest).JSON(errors.ErrInvalidID)
	}

	data, err := persistence.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
	}

	newTodo, err := data.DeleteTodo(id)

	if err != nil {
		if err.Error() == "todo not found" {
			return c.Status(fiber.StatusNotFound).JSON(errors.ErrTodoNotFound)
		}
		//file can not created
		return c.Status(fiber.StatusNotImplemented).JSON(errors.ErrFileNotCreated)
	}

	return c.JSON(newTodo)
}
