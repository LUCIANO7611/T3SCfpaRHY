// 代码生成时间: 2025-08-08 02:43:22
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// ErrorResponse is a struct to handle error responses
type ErrorResponse struct {
    Error string `json:"error"`
}

// User represents the user data model
type User struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// getUserByID is a handler function that retrieves a user by ID
func getUserByID(c *gin.Context) {
    id := c.Param("id")
    // Simulate database lookup
    user := User{ID: 1, Name: "John Doe", Email: "john.doe@example.com"}
    if id != "1" {
        c.JSON(http.StatusNotFound, ErrorResponse{Error: "User not found"})
    } else {
        c.JSON(http.StatusOK, user)
    }
}

// createUser is a handler function that creates a new user
func createUser(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
        return
    }
    // Simulate saving to database
    c.JSON(http.StatusCreated, newUser)
}

func main() {
    r := gin.Default()

    // Use Gin middleware
    r.Use(gin.Recovery())

    // Define routes
    r.GET("/user/:id", getUserByID)
    r.POST("/users", createUser)

    // Start the server
    r.Run(":8080")
}
