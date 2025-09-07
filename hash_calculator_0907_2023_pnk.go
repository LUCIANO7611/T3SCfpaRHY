// 代码生成时间: 2025-09-07 20:23:33
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// HashCalculator 结构体用于封装哈希计算的逻辑
type HashCalculator struct{}

// CalculateHash 是一个Gin处理器，用于计算字符串的SHA-256哈希值
func (h *HashCalculator) CalculateHash(c *gin.Context) {
    input := c.DefaultQuery("text", "")
    if input == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Missing 'text' parameter",
        })
        return
    }

    // 计算SHA-256哈希值
    hash := sha256.Sum256([]byte(input))
    hashString := hex.EncodeToString(hash[:])

    // 返回哈希值作为响应
    c.JSON(http.StatusOK, gin.H{
        "hash": hashString,
    })
}

func main() {
    r := gin.Default()

    // 实例化哈希计算工具
    hashCalculator := &HashCalculator{}

    // 路由设置
    r.GET("/hash", hashCalculator.CalculateHash)

    // 启动服务
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Error starting server: " + err.Error())
    }
}

// 注意：此代码假定运行环境已经安装了Gin框架