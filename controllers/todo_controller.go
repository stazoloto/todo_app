package controllers

import (
	"net/http"
	"strconv"
	"todo_app/models"

	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{Id: 1, Title: "Learn Go", Completed: false},
	{Id: 2, Title: "Build a Rest API", Completed: false},
}

// GetTodos - Получение всех задач
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// CreateTodo - Создание новой задачи
func CreateTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	newTodo.Id = len(todos) + 1
	todos = append(todos, newTodo)
	c.JSON(http.StatusCreated, newTodo)
}

// UpdateTodo - Обновление задачи
func UpdateTodo(c *gin.Context) {
	var updatedTodo models.Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		return
	}

	for i, t := range todos {
		if t.Id == updatedTodo.Id {
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

// DeleteTodo - удаление задачи
func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return
	}

	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}
