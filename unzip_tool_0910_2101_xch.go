// 代码生成时间: 2025-09-10 21:01:06
 * Features:
 * - Error handling
 * - Gin middleware usage (if needed)
 * - Follows Go best practices
 * - Includes comments and documentation
 */

package main

import (
    "archive/zip"
    "bytes"
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

// UnzipHandler is the Gin handler function for decompressing files.
func UnzipHandler(c *gin.Context) {
    // Check if file is uploaded
    if _, err := c.GetRawData(); err != nil || len(c.GetRawData()) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "No file uploaded",
        })
        return
    }

    // Create a buffer to hold the file data
    var buf bytes.Buffer
    _, err := io.Copy(&buf, c.Request.Body)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to read the file: %v", err),
        })
        return
    }
    defer c.Request.Body.Close()

    // Create a zip reader from the buffer
    zipReader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to create zip reader: %v", err),
        })
        return
    }

    // Extract files to system's temp directory
    tempDir, err := ioutil.TempDir("", "unzip-")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": fmt.Sprintf("Failed to create temp directory: %v", err),
        })
        return
    }
    defer os.RemoveAll(tempDir) // Clean up temp directory on function exit

    for _, zipFile := range zipReader.File {
        // Skip directories
        if zipFile.FileInfo().IsDir() {
            continue
        }

        // Create full path for each file
        filePath := filepath.Join(tempDir, zipFile.Name)
        if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to create directory: %v", err),
            })
            return
        }

        // Open file to write
        file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFile.Mode())
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to open file: %v", err),
            })
            return
        }
        defer file.Close()

        // Extract file
        fileInZip, err := zipFile.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to open zip file: %v", err),
            })
            return
        }
        defer fileInZip.Close()

        _, err = io.Copy(file, fileInZip)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": fmt.Sprintf("Failed to extract file: %v", err),
            })
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Files decompressed successfully",
        "temp_dir": tempDir,
    })
}

func main() {
    r := gin.Default()
    r.POST("/unzip", UnzipHandler)
    r.Run() // listening and serving on 0.0.0.0:8080
}
