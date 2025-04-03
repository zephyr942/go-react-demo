package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type Todo struct {
    ID          int    `json:"id"`
    Description string `json:"description"`
    IsComplete  bool   `json:"isComplete"`
}

var todos = []Todo{
    {ID: 1, Description: "Finish lecture", IsComplete: true},
    {ID: 2, Description: "Do homework", IsComplete: false},
    {ID: 3, Description: "Sleep", IsComplete: true},
}

var currentID = 4

func main() {
    r := gin.Default()

    // 允许跨域（前端 fetch 会用到）
    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // 获取所有任务
    r.GET("/todos", func(c *gin.Context) {
        c.JSON(http.StatusOK, todos)
    })

    // 添加新任务
    r.POST("/todos", func(c *gin.Context) {
        var newTodo Todo
        if err := c.ShouldBindJSON(&newTodo); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        newTodo.ID = currentID
        currentID++
        todos = append(todos, newTodo)
        c.JSON(http.StatusCreated, newTodo)
    })

    // 更新任务状态或描述
    r.PUT("/todos/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        var updateData Todo
        if err := c.ShouldBindJSON(&updateData); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        for i, todo := range todos {
            if todo.ID == id {
                todos[i].Description = updateData.Description
                todos[i].IsComplete = updateData.IsComplete
                c.JSON(http.StatusOK, todos[i])
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
    })

    // 删除任务
    r.DELETE("/todos/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }

        for i, todo := range todos {
            if todo.ID == id {
                todos = append(todos[:i], todos[i+1:]...)
                c.Status(http.StatusNoContent)
                return
            }
        }
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
    })

    r.Run(":8080") // 监听端口
}
