package Repository

import "awesomeProject/Entity"

type TodoPersistence interface {
	Todos() ([]Entity.Todo, error)
	TodoID(string) (Entity.Todo, error)
	CreateTodo(Entity.Todo) (Entity.Todo, error)
	UpdateTodo(string, Entity.Todo) (Entity.Todo, error)
}

type TodoRepo struct {
	Entity.Todo
}

func NewTodoPersistence() *TodoRepo {
	return &TodoRepo{Entity.Todo{}}
}

func (t *TodoRepo) Todos() ([]Entity.Todo, error) {
	return nil, nil
}

func (t *TodoRepo) TodoID(id string) (Entity.Todo, error) {
	var todo Entity.Todo

	return todo, nil
}

func (t *TodoRepo) CreateTodo(id string) (Entity.Todo, error) {
	var todo Entity.Todo

	return todo, nil
}

func (t *TodoRepo) UpdateTodo(todo Entity.Todo) (Entity.Todo, error) {

	return todo, nil
}
