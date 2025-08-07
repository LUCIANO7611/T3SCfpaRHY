// 代码生成时间: 2025-08-07 17:32:32
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "gopkg.in/go-playground/validator.v10"
)

// FormValidator defines the structure for the form data to be validated.
type FormValidator struct {
    Username string `form:"username" binding:"required,alphanum,min=3,max=30"`
    Email    string `form:"email" binding:"required,email"`
    Age      int    `form:"age" binding:"required,gte=18,lte=99"`
}

// ValidateForm is a Gin handler function that validates form data.
func ValidateForm(c *gin.Context) {
    var validator FormValidator
    // Use Gin's ShouldBind method to bind and validate the form data.
    if err := c.ShouldBind(&validator); err != nil {
        // Handle error with a JSON response.
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    } else {
        // If validation is successful, return a success message.
        c.JSON(http.StatusOK, gin.H{
            "message": "Form data is valid",
            "data": validator,
        })
    }
}

func main() {
    r := gin.Default()
    // Register the handler with the path "/form" and POST method.
    r.POST("/form", ValidateForm)
    // Use the Validator middleware to automatically validate
    // all requests based on their respective struct definitions.
    r.Use(ginrus.Ginrus())
    // Run the server.
    r.Run()
}
