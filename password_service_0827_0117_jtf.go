// 代码生成时间: 2025-08-27 01:17:07
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Key should be of AES key size (16, 24, or 32 bytes)
var key = []byte("your-very-secret-key")

// PasswordService struct containing necessary fields for encryption/decryption
type PasswordService struct{}

// EncryptHandler handles the encryption of the password
func (s *PasswordService) EncryptHandler(c *gin.Context) {
    var req struct {
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{'error': 'Invalid request'})
        return
    }
    encryptedPassword, err := encrypt(req.Password, key)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{'error': 'Encryption failed'})
        return
    }
    c.JSON(http.StatusOK, gin.H{'encryptedPassword': encryptedPassword})
}

// DecryptHandler handles the decryption of the password
func (s *PasswordService) DecryptHandler(c *gin.Context) {
    var req struct {
        EncryptedPassword string `json:"encryptedPassword"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{'error': 'Invalid request'})
        return
    }
    decryptedPassword, err := decrypt(req.EncryptedPassword, key)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{'error': 'Decryption failed'})
        return
    }
    c.JSON(http.StatusOK, gin.H{'decryptedPassword': decryptedPassword})
}

// Main function to setup and run the Gin server
func main() {
    r := gin.Default()
    ps := &PasswordService{}
    r.POST("/encrypt", ps.EncryptHandler)
    r.POST("/decrypt", ps.DecryptHandler)
    r.Run() // listen and serve on 0.0.0.0:8080
}

// encrypt function for AES encryption
func encrypt(plaintext string, key []byte) (string, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    
    plaintextBytes := []byte(plaintext)
    
    // PKCS#7 padding
    padding := aes.BlockSize - len(plaintextBytes)%aes.BlockSize
    plaintextBytes = append(plaintextBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)
    
    encryptedBytes := make([]byte, aes.BlockSize+len(plaintextBytes))
    iv := encryptedBytes[:aes.BlockSize]
    
    if _, err := rand.Read(iv); err != nil {
        return "", err
    }
    
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(encryptedBytes[aes.BlockSize:], plaintextBytes)
    
    return base64.StdEncoding.EncodeToString(encryptedBytes), nil
}

// decrypt function for AES decryption
func decrypt(encrypted string, key []byte) (string, error) {
    encryptedBytes, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil {
        return "", err
    }
    
    if len(encryptedBytes) < aes.BlockSize {
        return "", fmt.Errorf("ciphertext too short")
    }
    
    iv := encryptedBytes[:aes.BlockSize]
    encryptedBytes = encryptedBytes[aes.BlockSize:]
    
    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }
    
    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(encryptedBytes, encryptedBytes)
    
    // Remove padding
    padding := byte(encryptedBytes[len(encryptedBytes)-1])
    encryptedBytes = encryptedBytes[:len(encryptedBytes)-int(padding)]
    
    return string(encryptedBytes), nil
}