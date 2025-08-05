// 代码生成时间: 2025-08-06 07:16:54
package main

import (
    "net"
    "time"
    "github.com/gin-gonic/gin"
)

// NetworkStatusChecker 结构体用于存储网络检查的配置
type NetworkStatusChecker struct {
    // 可以添加配置项，比如超时时间、检查频率等
    timeout time.Duration
}

// NewNetworkStatusChecker 创建一个新的NetworkStatusChecker实例
func NewNetworkStatusChecker(timeout time.Duration) *NetworkStatusChecker {
    return &NetworkStatusChecker{
        timeout: timeout,
    }
}

// CheckNetworkStatus 实现网络连接状态检查
func (c *NetworkStatusChecker) CheckNetworkStatus(host string) (bool, error) {
    // 创建一个TCP连接
    conn, err := net.DialTimeout("tcp", host, c.timeout)
    if err != nil {
        return false, err
    }
    defer conn.Close()
    return true, nil
}

func main() {
    r := gin.Default()

    // 创建NetworkStatusChecker实例
    nsChecker := NewNetworkStatusChecker(5 * time.Second)

    // 定义GET路由，用于检查网络状态
    r.GET("/check", func(c *gin.Context) {
        host := c.Query("host")
        if host == "" {
            c.JSON(400, gin.H{
                "error": "Host parameter is required",
            })
            return
        }

        // 检查网络状态
        isReachable, err := nsChecker.CheckNetworkStatus(host)
        if err != nil {
            c.JSON(500, gin.H{
                "error": "Failed to check network status",
            })
            return
        }

        // 返回检查结果
        c.JSON(200, gin.H{
            "host": host,
            "is_reachable": isReachable,
        })
    })

    // 启动Gin服务
    r.Run() // 默认在8080端口启动
}
