// 代码生成时间: 2025-09-06 00:22:09
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "errors"
)

// FormValidator 是一个表单验证的结构体
type FormValidator struct {
    Username string `form:"username" binding:"required,alphanum"`
    Email    string `form:"email" binding:"required,email"`
}

// ValidateForm 处理表单验证的中间件
func ValidateForm(c *gin.Context) {
    var validator FormValidator
    if err := c.ShouldBind(&validator); err != nil {
        // 错误处理
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        c.Abort()
        return
    }
    // 继续处理请求
    c.Next()
}

func main() {
    r := gin.Default()

    // 使用表单验证中间件
    r.POST("/form", ValidateForm, func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Form is valid",
        })
    })

    // 启动服务器
    r.Run()
}
