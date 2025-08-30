// 代码生成时间: 2025-08-30 17:56:42
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// FolderOrganizer 结构体用于存储文件夹信息
type FolderOrganizer struct {
    Source string
    Target string
}

// Handler 处理文件夹整理请求
func (f *FolderOrganizer) Handler(c *gin.Context) {
    // 解析请求参数
    source := c.PostForm("source")
    target := c.PostForm("target")

    // 检查参数是否为空
    if source == "" || target == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "source and target parameters are required",
        })
        return
    }

    // 创建FolderOrganizer实例
    organizer := FolderOrganizer{
        Source: source,
        Target: target,
    }

    // 调用整理函数
    err := organizer.organize()
    if err != nil {
        // 错误处理
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    // 返回成功响应
    c.JSON(http.StatusOK, gin.H{
        "message": "Folder structure organized successfully",
    })
}

// organize 实现文件夹整理逻辑
func (f *FolderOrganizer) organize() error {
    // 读取源文件夹内容
    files, err := os.ReadDir(f.Source)
    if err != nil {
        return fmt.Errorf("failed to read source folder: %w", err)
    }

    // 创建目标文件夹
    err = os.MkdirAll(f.Target, os.ModePerm)
    if err != nil {
        return fmt.Errorf("failed to create target folder: %w", err)
    }

    // 遍历源文件夹文件并复制到目标文件夹
    for _, file := range files {
        srcPath := filepath.Join(f.Source, file.Name())
        dstPath := filepath.Join(f.Target, file.Name())

        // 检查是否为文件夹，如果是则递归调用organize
        if file.IsDir() {
            err = (&FolderOrganizer{
                Source: srcPath,
                Target: dstPath,
            }).organize()
            if err != nil {
                return fmt.Errorf("failed to organize subfolder %s: %w", file.Name(), err)
            }
        } else {
            // 复制文件
            err = copyFile(srcPath, dstPath)
            if err != nil {
                return fmt.Errorf("failed to copy file %s: %w", file.Name(), err)
            }
        }
    }

    return nil
}

// copyFile 复制文件
func copyFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source file: %w", err)
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    if err != nil {
        return fmt.Errorf("failed to copy file: %w", err)
    }

    return nil
}

func main() {
    r := gin.Default()

    // 注册处理器
    r.POST("/organize", (&FolderOrganizer{}).Handler)

    // 启动服务器
    log.Fatal(r.Run(":8080"))
}
