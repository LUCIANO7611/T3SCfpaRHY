// 代码生成时间: 2025-08-13 02:21:08
package main

import (
    "fmt"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gin-gonic/gin"
)

// BatchRenameRequest 结构体定义了批量重命名的请求数据
type BatchRenameRequest struct {
    SourceFolder string   `json:"sourceFolder"`
    Files       []string `json:"files"`
    NewNames    []string `json:"newNames"`
}

// BatchRenameResponse 结构体定义了批量重命名的响应数据
type BatchRenameResponse struct {
    Successes []string `json:"successes"`
    Errors   []string `json: