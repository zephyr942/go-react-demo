package routes

import (
    "github.com/zephyr942/backend-go/controllers"

    "github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(r *gin.Engine) {
    todo := r.Group("/api/todos")
    {
        todo.GET("", controllers.GetTodos)
        todo.POST("", controllers.CreateTodo)
        todo.PUT("/:id", controllers.UpdateTodo)
        todo.DELETE("/:id", controllers.DeleteTodo)
    }
}
