package http

import (
	"net/http"
	"strconv"
	"todo_app/internal/app"
	"todo_app/internal/domain"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	todoUsecase app.TodoUsecase
}

func NewTodoController(u app.TodoUsecase) *TodoController {
	return &TodoController{todoUsecase: u}
}

func(ctrl *TodoController) GetTodos(c *gin.Context) {
	todos, err := ctrl.todoUsecase.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

func(ctrl *TodoController) CreateTodo(c *gin.Context) {
	var todo domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.todoUsecase.CreateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func(ctrl *TodoController) UpdateTodo(c *gin.Context) {
	var todo domain.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.todoUsecase.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

func(ctrl *TodoController) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	if err := ctrl.todoUsecase.DeleteTodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}