package domain

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// Интерфейс для репозиториев
type TodoRepository interface {
	GetAll() ([]Todo, error)
	Save(todo Todo) error
	Update(todo Todo) error
	Delete(id int) error
}