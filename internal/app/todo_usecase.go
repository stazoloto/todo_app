package app

import "todo_app/internal/domain"

type TodoUsecase interface {
	GetTodos() ([]domain.Todo, error)
	CreateTodo(todo domain.Todo) error
	UpdateTodo(todo domain.Todo) error
	DeleteTodo(id int) error
}

type todoUsecase struct {
	todoRepo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepo: repo}
}

func(u *todoUsecase) GetTodos() ([]domain.Todo, error) {
	return u.todoRepo.GetAll()
}

func(u *todoUsecase) CreateTodo(todo domain.Todo) error {
	return u.todoRepo.Save(todo)
}

func(u *todoUsecase) UpdateTodo(todo domain.Todo) error {
	return u.todoRepo.Update(todo)
}

func(u *todoUsecase) DeleteTodo(id int) error {
	return u.todoRepo.Delete(id)
}
