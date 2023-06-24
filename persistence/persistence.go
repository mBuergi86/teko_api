package persistence

import (
	"awesomeProject/entity"
	"awesomeProject/utility"
	"errors"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type TodoPersistence struct {
	todos []entity.Todo
}

func NewTodoPersistence() (*TodoPersistence, error) {
	file, err := utility.OpenRead()
	if err != nil {
		return nil, fiber.ErrNotFound
	}

	var todos []entity.Todo
	for _, csvData := range file {
		todos = append(todos, csvData.Todo...)
	}

	return &TodoPersistence{todos}, nil
}

func (t *TodoPersistence) Todos() []entity.Todo {
	return t.todos
}

func (t *TodoPersistence) TodoID(id string) (entity.Todo, error) {
	for _, todo := range t.todos {
		if todo.Id == id {
			return todo, nil
		}
	}
	return entity.Todo{}, errors.New("todo not found")
}

func (t *TodoPersistence) CreateTodo(newTask entity.Todo) entity.Todo {
	nextID := utility.NextID(t.todos)
	newTask.Id = strconv.Itoa(nextID)

	err := utility.WriteSaveFile(entity.CSVData{Todo: []entity.Todo{newTask}})
	if err != nil {
		return entity.Todo{}
	}

	return newTask
}

func (t *TodoPersistence) UpdateTodo(id string, updateTask entity.Todo) (entity.Todo, error) {
	newTodos := make([]entity.Todo, 0)

	for i, item := range t.todos {
		if item.Id == id {
			t.todos[i] = updateTask
			newTodos = append(newTodos, t.todos...)

			err := utility.WriteUpdateFile(entity.CSVData{Todo: newTodos})
			if err != nil {
				return entity.Todo{}, errors.New("file can't created")
			}
			return updateTask, nil
		}
	}

	return entity.Todo{}, errors.New("todo not found")
}

func (t *TodoPersistence) DeleteTodo(id string) ([]entity.Todo, error) {
	for i, item := range t.todos {
		if item.Id == id {
			t.todos = append(t.todos[:i], t.todos[i+1:]...)

			err := utility.WriteUpdateFile(entity.CSVData{Todo: t.todos})
			if err != nil {
				return nil, errors.New("file can't be created")
			}

			return t.todos, nil
		}
	}

	return nil, errors.New("todo not found")
}
