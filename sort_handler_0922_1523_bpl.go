// 代码生成时间: 2025-09-22 15:23:00
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// SortHandler is a Gin handler function for sorting algorithms.
// It will receive a request with a list of integers and return the sorted list.
func SortHandler(c *gin.Context) {
    // Define the structure for the incoming request data.
    type RequestData struct {
        Numbers []int `json:"numbers" binding:"required,min=1"`
    }

    // Create an instance of RequestData to bind the request data to.
    var requestData RequestData

    // Bind the JSON data from the request to requestData.
    if err := c.ShouldBindJSON(&requestData); err != nil {
        // Return a bad request error if binding fails.
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid request data."
        })
        return
    }

    // Sort the numbers using a simple bubble sort algorithm.
    // This is just for demonstration purposes; in production, use a more efficient algorithm.
    for i := 0; i < len(requestData.Numbers); i++ {
        for j := 0; j < len(requestData.Numbers)-1-i; j++ {
            if requestData.Numbers[j] > requestData.Numbers[j+1] {
                requestData.Numbers[j], requestData.Numbers[j+1] = requestData.Numbers[j+1], requestData.Numbers[j]
            }
        }
    }

    // Return the sorted list as a JSON response.
    c.JSON(http.StatusOK, gin.H{
        "sorted_numbers": requestData.Numbers,
    })
}

func main() {
    // Create a new Gin router.
    router := gin.Default()

    // Add the SortHandler to the router with the path "/sort".
    router.POST("sort", SortHandler)

    // Start the server on port 8080.
    router.Run(":8080")
}
