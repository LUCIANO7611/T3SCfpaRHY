// 代码生成时间: 2025-07-30 18:00:15
package main

import (
    "encoding/json"
    "net/http"
    "github.com/gin-gonic/gin"
    "log"
    "github.com/xuri/excelize/v2"
)

// ExcelGenerator 结构体，用于生成Excel文档
type ExcelGenerator struct {
    // 标题行
    Title []string
    // 数据行
    Data [][]string
}

// NewExcelGenerator 创建一个ExcelGenerator实例
func NewExcelGenerator(title []string, data [][]string) *ExcelGenerator {
    return &ExcelGenerator{
        Title: title,
        Data: data,
    }
}

// GenerateExcel 生成Excel文件并返回文件名
func (eg *ExcelGenerator) GenerateExcel() (string, error)
{
    f := excelize.NewFile()
    // 创建标题行
    titleRow := "A1"
    for i, title := range eg.Title {
        f.SetCellValue("Sheet1", titleRow, i+1, title)
        titleRow = IncrementCell(titleRow)
    }
    // 创建数据行
    for i, row := range eg.Data {
        dataRow := fmt.Sprintf("A%d", i+2)
        for j, value := range row {
            f.SetCellValue("Sheet1", dataRow, j+1, value)
            dataRow = IncrementCell(dataRow)
        }
    }
    // 保存文件
    filename := "example.xlsx"
    if err := f.SaveAs(filename); err != nil {
        return "", err
    }
    return filename, nil
}

// IncrementCell 递增Excel单元格位置
func IncrementCell(cell string) string {
    // 此函数用于递增Excel单元格位置，例如A1变为A2
    // 此处省略实现细节
    return cell
}

func main() {
    r := gin.Default()
    r.Use(gin.Recovery()) // 使用Recovery中间件以处理panic
    
    // 定义路由和处理器
    r.GET("/generate", func(c *gin.Context) {
        generator := NewExcelGenerator([]string{"Name", "Age"}, [][]string{{"Alice", "30"}, {"Bob", "25"}})
        filename, err := generator.GenerateExcel()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "Failed to generate Excel file",
            })
            return
        }
        c.JSON(http.StatusOK, gin.H{
            "filename": filename,
            "message": "Excel file generated successfully",
        })
    })
    
    // 启动服务器
    r.Run()
}
