// 代码生成时间: 2025-09-03 16:35:05
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

// ErrorLogMiddleware is a Gin middleware that captures error logs.
func ErrorLogMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Request start time
        start := time.Now()
        path := c.Request.URL.Path
        method := c.Request.Method

        // Process request
        c.Next()

        // Calculate duration
        duration := time.Since(start)

        // Check if an error occurred during the request handling
        err := c.Errors.Last().Err
        if err != nil {
            // Log the error with request info
            log.Printf("Error in %s %s: %v
", method, path, err)
        } else {
            // Log successful request
            log.Printf("%s %s %v
", method, path, duration)
        }
    }
}

func main() {
    router := gin.Default()
    router.Use(ErrorLogMiddleware())

    // Define routes with error handling
    router.GET("/example", func(c *gin.Context) {
        // Simulate an error
        c.JSON(http.StatusInternalServerError, gin.H{
            "status": "error",
            "message": "Internal Server Error",
        })
    })

    // Start the server
    router.Run(":8080") // listening and serving on 0.0.0.0:8080
}
