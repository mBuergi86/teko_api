// Package models contains the todo backend models and corresponding business logic
package models

import (
	"errors"
	"sort"
	"strconv"
)

type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Terminated  bool   `json:"terminated"`
}

func (t Todo) Serialize() []string {
	todoSerialized := []string{t.Id, t.Title, t.Description, strconv.FormatBool(t.Terminated)}
	return todoSerialized
}

type JsonExtendedResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonDataResponse struct {
	Data []Todo `json:"data"`
}

type JsonErrorResponse struct {
	Error ApiError `json:"error"`
}

type ApiError struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

var todoPersistence TodoPersistence

// SetTodoPersistence allows to set the persistence type
func SetTodoPersistence(todoPersistenceNew TodoPersistence) error {
	if todoPersistenceNew == nil {
		return errors.New("todo persistence must not be nil")
	}
	todoPersistence = todoPersistenceNew
	return nil
}

// ReadTodos returns todo's from persistence layer
func ReadTodos() ([]Todo, error) {
	if todoPersistence == nil {
		return nil, errors.New("todo persistence must not be nil")
	}
	return todoPersistence.ReadTodos()
}

// ReadTodoById returns todo with passed id when existing from persistence layer
func ReadTodoById(id string) (Todo, error) {
	if todoPersistence == nil {
		return Todo{}, errors.New("todo persistence must not be nil")
	}
	return todoPersistence.ReadTodoById(id)
}

// CreateTodo stores the passed todo in the persistence layer and returns the stored todo
func CreateTodo(todo Todo) (Todo, error) {
	if todoPersistence == nil {
		return Todo{}, errors.New("todo persistence must not be nil")
	}
	return todoPersistence.CreateTodo(todo)
}

func UpdateTodo(id string, todo Todo) (Todo, error) {
	if todoPersistence == nil {
		return Todo{}, errors.New("todo persistence must not be nil")
	}
	return todoPersistence.UpdateTodo(id, todo)
}

// SortTodosAfterIdAscending sorts the todos ascending after the id and returns sorted todos
func SortTodosAfterIdAscending(todos []Todo) []Todo {
	sort.Slice(todos, func(i, j int) bool {
		leftValueAsInt, _ := strconv.Atoi(todos[i].Id)
		rightValueAsInt, _ := strconv.Atoi(todos[j].Id)
		return leftValueAsInt < rightValueAsInt
	})
	return todos
}
