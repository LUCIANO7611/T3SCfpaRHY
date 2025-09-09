// 代码生成时间: 2025-09-09 12:06:34
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/hex"
    "errors"
    "github.com/gin-gonic/gin"
    "log"
)

// PasswordService represents the service for password encryption and decryption.
type PasswordService struct {
    // Your encryption key, must be 16, 24, or 32 bytes long.
    key []byte
}

// NewPasswordService creates a new instance of PasswordService.
func NewPasswordService(key string) *PasswordService {
    return &PasswordService{
        key: []byte(key),
    }
}

// Encrypt encrypts the given password using AES encryption.
func (ps *PasswordService) Encrypt(password string) (string, error) {
    if len(password)%aes.BlockSize != 0 {
        return "", errors.New("password size must be a multiple of the block size")
    }

    block, err := aes.NewCipher(ps.key)
    if err != nil {
        return "", err
    }

    // The IV needs to be unique but not secure.
    // Commonly, it's called nonce - number used once.
    // For this example, we use a fixed array. In production,
    // you would use a cryptographically secure random number generator.
    iv := []byte("this is an iv456")
    if len(iv) != aes.BlockSize {
        panic("iv size must be equal to the block size")
    }

    encrypted := make([]byte, aes.BlockSize+len(password))
    copy(encrypted[aes.BlockSize:], []byte(password))

    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(encrypted[aes.BlockSize:], []byte(password))

    // It's important to ensure that the base64 encoded string does not contain any carriage returns.
    return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts the encrypted password using AES decryption.
func (ps *PasswordService) Decrypt(encrypted string) (string, error) {
    encryptedData, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil {
        return "", err
    }

    if len(encryptedData)%aes.BlockSize != 0 {
        return "", errors.New("ciphertext size must be a multiple of the block size")
    }

    block, err := aes.NewCipher(ps.key)
    if err != nil {
        return "", err
    }

    iv := []byte("this is an iv456")
    if len(iv) != aes.BlockSize {
        panic("iv size must be equal to the block size")
    }

    decrypted := make([]byte, len(encryptedData))
    mode := cipher.NewCBCDecrypter(block, iv)
    mode.CryptBlocks(decrypted, encryptedData)

    return hex.EncodeToString(decrypted[aes.BlockSize:]), nil
}

func main() {
    // Create a new instance of PasswordService with a predefined key.
    ps := NewPasswordService("your-256-bit-encryption-key")

    // Setup the Gin router.
    router := gin.Default()

    router.POST("/encrypt", func(c *gin.Context) {
        password := c.PostForm("password")
        encrypted, err := ps.Encrypt(password)
        if err != nil {
            c.JSON(500, gin.H{
                "error": err.Error(),
            })
        } else {
            c.JSON(200, gin.H{
                "encrypted": encrypted,
            })
        }
    })

    router.POST("/decrypt", func(c *gin.Context) {
        encrypted := c.PostForm("encrypted")
        decrypted, err := ps.Decrypt(encrypted)
        if err != nil {
            c.JSON(500, gin.H{
                "error": err.Error(),
            })
        } else {
            c.JSON(200, gin.H{
                "decrypted": decrypted,
            })
        }
    })

    // Start the server on port 8080.
    log.Fatal(router.Run(":8080"))
}
