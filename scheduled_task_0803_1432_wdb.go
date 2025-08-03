// 代码生成时间: 2025-08-03 14:32:29
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

// Task represents a task to be scheduled
type Task struct {
    Name    string
    Execute func() error
}

// TaskScheduler is a struct to hold the scheduled tasks
type TaskScheduler struct {
    tasks map[string]Task
}

// NewTaskScheduler creates a new instance of TaskScheduler
func NewTaskScheduler() *TaskScheduler {
    return &TaskScheduler{
        tasks: make(map[string]Task),
    }
}

// AddTask adds a task to the scheduler
func (s *TaskScheduler) AddTask(name string, task Task) {
    s.tasks[name] = task
}

// Run runs all scheduled tasks
func (s *TaskScheduler) Run() {
    for name, task := range s.tasks {
        go func(name string, task Task) {
            if err := task.Execute(); err != nil {
                log.Printf("Error executing task %s: %v", name, err)
            }
        }(name, task)
    }
}

func main() {
    r := gin.Default()
    // Middleware for logging requests
    r.Use(gin.Logger())
    // Middleware for recovery from panic
    r.Use(gin.Recovery())

    scheduler := NewTaskScheduler()

    // Example task: Print a message every 5 seconds
    scheduler.AddTask("exampleTask", Task{
        Name: "Example Task",
        Execute: func() error {
            fmt.Println("Task Executed")
            return nil
        },
    })

    // Endpoint to trigger the tasks
    r.GET("/run-tasks", func(c *gin.Context) {
        if err := scheduler.Run(); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to run tasks",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "message": "Tasks executed",
        })
    })

    // Run the Gin server
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}