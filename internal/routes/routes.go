package routes

import (
	"github.com/gin-gonic/gin"
	"todo_app/internal/adapter/http"
)

func SetupRouter(todoController *http.TodoController) *gin.Engine {
	r := gin.Default()

	r.GET("/todos", todoController.GetTodos)
	r.POST("/todos", todoController.CreateTodo)
	r.PUT("/todos", todoController.UpdateTodo)
	r.DELETE("/todos/:id", todoController.DeleteTodo)

	return r
}
