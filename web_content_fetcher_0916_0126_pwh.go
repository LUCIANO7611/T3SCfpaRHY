// 代码生成时间: 2025-09-16 01:26:07
package main

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/PuerkitoBio/goquery"
)

// WebContentFetcherHandler 是抓取网页内容的处理器
func WebContentFetcherHandler(c *gin.Context) {
    url := c.Query("url")
    if url == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "url parameter is required",
        })
        return
    }

    response, err := http.Get(url)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to fetch URL: %v", err),
        })
        return
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to fetch URL: Status Code %d", response.StatusCode),
        })
        return
    }

    doc, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to parse document: %v", err),
        })
        return
    }

    // 抓取网页的HTML内容
    html, err := doc.Html()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to get HTML: %v", err),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "url": url,
        "html": html,
    })
}

func main() {
    r := gin.Default()
    r.GET("/fetch", WebContentFetcherHandler)
    r.Run(":8080") // 默认在8080端口启动服务
}
