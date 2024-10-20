package main

import (
	"todo_app/internal/adapter/db"
	"todo_app/internal/adapter/http"
	"todo_app/internal/app"
	"todo_app/internal/routes"
)

func main() {
	// Инициализация репозитория и Use Case
	todoRepo := db.NewTodoRepository()
	todoUsecase := app.NewTodoUsecase(todoRepo)
	todoController := http.NewTodoController(todoUsecase)

	// Настройка роутов
	r := routes.SetupRouter(todoController)
	r.Run(":8080")
}
