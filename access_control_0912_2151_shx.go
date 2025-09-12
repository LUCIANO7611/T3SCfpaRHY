// 代码生成时间: 2025-09-12 21:51:11
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Middleware for authentication
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract the token from the header
        token := c.Request.Header.Get("Authorization")
        if token == "" {
            // If no token is provided, return an unauthorized error
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Authorization token is missing",
            })
            c.Abort()
            return
        }
        // Here you would validate the token against your auth service
        // For this example, we'll assume any non-empty token is valid
        c.Next()
    }
}

// ErrorHandler middleware for handling errors
func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        if len(c.Errors) > 0 {
            // If there are any errors, return the first one
            err := c.Errors.Last().Err
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": err.Error(),
            })
        }
    }
}

func main() {
    router := gin.Default()

    // Register the authentication middleware
    router.Use(AuthMiddleware())
    // Register the error handler middleware
    router.Use(ErrorHandler())

    // Define a protected route that requires authentication
    router.GET("/protected", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "This is a protected route",
        })
    })

    // Start the server
    router.Run(":8080")
}
