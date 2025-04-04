package controllers

import (
    "net/http"
    "github.com/zephyr942/backend-go/models"
    "github.com/zephyr942/backend-go/services"

    "github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    todos, err := services.GetAllTodos()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.CreateTodo(todo)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusCreated)
}

func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var updated models.Todo
    if err := c.ShouldBindJSON(&updated); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.UpdateTodo(id, updated)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusOK)
}

func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    err := services.DeleteTodo(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
