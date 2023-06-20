// Package models contains the todo backend models and corresponding business logic
package models

type TodoPersistence interface {
	ReadTodos() ([]Todo, error)
	ReadTodoById(string) (Todo, error)
	CreateTodo(Todo) (Todo, error)
	UpdateTodo(string, Todo) (Todo, error)
}
