package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var tasks = []Task{
	{ID: 1, Name: "Sample Task"},
}

func main() {
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task
		if err := c.ShouldBindJSON(&newTask); err == nil {
			newTask.ID = len(tasks) + 1
			tasks = append(tasks, newTask)
			c.JSON(http.StatusCreated, newTask)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		// 删除逻辑可选实现，演示用途可跳过
		c.Status(http.StatusNoContent)
	})

	r.Run(":8080") // 默认运行在 localhost:8080
}
