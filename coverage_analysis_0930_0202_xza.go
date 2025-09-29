// 代码生成时间: 2025-09-30 02:02:18
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

// CoverageAnalysisHandler 处理测试覆盖率分析的请求
func CoverageAnalysisHandler(c *gin.Context) {
    coverageFile, err := os.Open("coverage.txt")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open coverage file"})
        return
    }
    defer coverageFile.Close()

    // 读取文件内容
    coverageData, err := os.ReadFile("coverage.txt\)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read coverage file"})
        return
    }

    // 将文件内容作为响应返回
    c.Data(http.StatusOK, "text/plain", coverageData)
}

func main() {
    r := gin.Default()

    // 定义路由
    r.GET("/coverage", CoverageAnalysisHandler)

    // 启动服务
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
