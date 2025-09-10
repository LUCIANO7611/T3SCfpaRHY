// 代码生成时间: 2025-09-10 08:12:29
package main

import (
    "net/http"
    "net/url"
    "strings"

    "github.com/gin-gonic/gin"
)

// URLValidator 验证URL链接是否有效
type URLValidator struct {
    // 可添加其他需要的字段
}

// ValidateURL 验证给定的URL是否有效
func (validator *URLValidator) ValidateURL(c *gin.Context) {
    rawURL := c.Query("url")
    if rawURL == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL parameter is required",
        })
        return
    }

    // 尝试解析URL
    parsedURL, err := url.ParseRequestURI(rawURL)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid URL format",
        })
        return
    }

    // 检查URL是否包含协议（scheme）
    if strings.Contains(parsedURL.Scheme, "://") {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "URL scheme is not allowed",
        })
        return
    }

    // 这里可以添加更多URL验证逻辑，例如检查主机名是否有效等

    c.JSON(http.StatusOK, gin.H{
        "message": "URL is valid",
    })
}

func main() {
    r := gin.Default()

    // 注册URL验证处理器
    r.GET("/check_url", func(c *gin.Context) {
        validator := &URLValidator{}
        validator.ValidateURL(c)
    })

    // 启动服务
    r.Run()
}
