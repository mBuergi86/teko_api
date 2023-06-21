package handler

import (
	"awesomeProject/entity"
	"awesomeProject/repository"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Todos(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, err := repository.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	return c.JSON(data.Todos())
}

func TodoID(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	data, err := repository.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	todo, err := data.TodoID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	data, err := repository.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var todo entity.Todo

	if err := json.Unmarshal(c.Body(), &todo); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(data.CreateTodo(todo))
}

func UpdateTodo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	data, err := repository.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	var todo entity.Todo

	if err := json.Unmarshal(c.Body(), &todo); err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	newTodo, err := data.UpdateTodo(id, todo)

	if err != nil {
		if err.Error() == "todo not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		//file can not created
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(newTodo)
}

func DeleteTodo(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "Invalid ID",
		})
	}

	data, err := repository.NewTodoPersistence()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	newTodo, err := data.DeleteTodo(id)

	if err != nil {
		if err.Error() == "todo not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		//file can not created
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(newTodo)
}
