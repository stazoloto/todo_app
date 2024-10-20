package db

import "todo_app/internal/domain"

type todoRepository struct {
	todos []domain.Todo
}

func NewTodoRepository() domain.TodoRepository {
	return &todoRepository{
		todos: []domain.Todo{
			{Id: 0, Title: "Learn Go", Completed: false},
			{Id: 1, Title: "Build a REST API", Completed: false},
		},
	}
}

func (r *todoRepository) GetAll() ([]domain.Todo, error) {
	return r.todos, nil
}

func (r *todoRepository) Save(todo domain.Todo) error {
	todo.Id = len(r.todos) + 1
	r.todos = append(r.todos, todo)
	return nil
}

func (r *todoRepository) Update(todo domain.Todo) error {
	for i, t := range r.todos {
		if t.Id == todo.Id {
			r.todos[i] = todo
			return nil
		}
	}
	return nil
}

func (r *todoRepository) Delete(id int) error {
	for i, t := range r.todos {
		if t.Id == id {
			r.todos = append(r.todos[:i], r.todos[i+1:]...)
			return nil
		}
	}
	return nil
}
