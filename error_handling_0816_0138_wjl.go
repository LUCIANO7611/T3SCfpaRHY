// 代码生成时间: 2025-08-16 01:38:49
package main

import (
    "fmt"
    "net/http"
    "gin-gonic/gin"
)

// ErrorResponse 用于定义API错误响应的结构
type ErrorResponse struct {
    Error string `json:"error"`
}

// ErrorResponseHandler 创建一个错误响应中间件
func ErrorResponseHandler(c *gin.Context) {
    c.Next()
    if len(c.Errors.ByType(gin.ErrorTypeAny)) > 0 {
        // 如果有错误发生，构造ErrorResponse
        err := c.Errors.Last().Err
        response := ErrorResponse{Error: fmt.Sprintf("%v", err)}
        // 设置响应状态码和错误信息
        c.AbortWithStatusJSON(http.StatusInternalServerError, response)
    }
}

// ExampleHandler 是一个示例处理器，演示如何使用中间件和错误处理
func ExampleHandler(c *gin.Context) {
    // 模拟一个错误
    c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal Server Error"})
}

func main() {
    r := gin.Default()
    r.Use(ErrorResponseHandler)

    // 注册处理器
    r.GET("/example", ExampleHandler)

    // 启动服务器
    r.Run()
}
