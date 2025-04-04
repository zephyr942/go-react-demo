package main

import (
    "github.com/zephyr942/backend-go/database"
    "github.com/zephyr942/backend-go/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect() // Connect to MongoDB

    r := gin.Default()

    // Enable CORS for frontend access
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

    routes.RegisterTodoRoutes(r)

    r.Run(":8080") // Start server
}
