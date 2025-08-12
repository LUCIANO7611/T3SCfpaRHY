// 代码生成时间: 2025-08-12 11:20:58
package main

import (
    "net/http"
    "regexp"
    "strings"
    "github.com/gin-gonic/gin"
)

// URLValidationPattern is a regular expression for validating URLs.
const URLValidationPattern = `^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`

// validateURL checks if the provided URL is valid.
func validateURL(c *gin.Context) {
    url := c.Query("url")
    if url == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL parameter is required",
        })
        return
    }

    re := regexp.MustCompile(URLValidationPattern)
    if !re.MatchString(url) {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid URL format",
        })
        return
    }

    // Additional URL validation logic can be added here.
    // For example, check if the URL is reachable, etc.
    // ...

    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid",
    })
}

func main() {
    r := gin.Default()

    // Add middleware if needed.
    // For example, logging middleware:
    // r.Use(gin.Logger())
    // Recovery middleware:
    // r.Use(gin.Recovery())

    r.GET("/check", validateURL)

    // Start the server on port 8080
    r.Run()
}
