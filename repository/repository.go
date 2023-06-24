package repository

import (
	"awesomeProject/entity"
)

type TodoRepository interface {
	Todos() []entity.Todo
	TodoID(string) (entity.Todo, error)
	CreateTodo(entity.Todo) entity.Todo
	UpdateTodo(string, entity.Todo) (entity.Todo, error)
	DeleteTodo(string) ([]entity.Todo, error)
}
