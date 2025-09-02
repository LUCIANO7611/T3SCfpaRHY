// 代码生成时间: 2025-09-02 15:59:31
package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

// ShoppingCart represents a simple shopping cart with basic operations.
type ShoppingCart struct {
    Items map[string]int
}

// NewShoppingCart creates a new instance of ShoppingCart.
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        Items: make(map[string]int),
    }
}

// AddItem adds an item to the cart with a specified quantity.
func (cart *ShoppingCart) AddItem(itemID string, quantity int) {
    cart.Items[itemID] += quantity
}

// RemoveItem removes an item from the cart with a specified quantity.
func (cart *ShoppingCart) RemoveItem(itemID string, quantity int) error {
    if quantity <= 0 {
        return fmt.Errorf("quantity must be greater than zero")
    }
    if _, exists := cart.Items[itemID]; !exists {
        return fmt.Errorf("item '%s' does not exist in the cart", itemID)
    }
    cart.Items[itemID] -= quantity
    if cart.Items[itemID] <= 0 {
        delete(cart.Items, itemID)
    }
    return nil
}

// GetCart returns the current state of the cart.
func (cart *ShoppingCart) GetCart() map[string]int {
    return cart.Items
}

func main() {
    router := gin.Default()

    // middleware for logging requests
    router.Use(func(c *gin.Context) {
        fmt.Printf("[INFO] %s %s %s
", c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery)
        c.Next()
    })

    // Middleware for recovering from any panics and setting HTTP 500 status code
    router.Use(gin.Recovery())

    cart := NewShoppingCart()

    // POST endpoint to add item to the cart
    router.POST("/add", func(c *gin.Context) {
        itemID := c.PostForm("itemID")
        quantity, err := strconv.Atoi(c.PostForm("quantity"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid quantity",
            })
            return
        }

        cart.AddItem(itemID, quantity)

        c.JSON(http.StatusOK, cart.GetCart())
    })

    // POST endpoint to remove item from the cart
    router.POST("/remove", func(c *gin.Context) {
        itemID := c.PostForm("itemID")
        quantity, err := strconv.Atoi(c.PostForm("quantity"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Invalid quantity",
            })
            return
        }

        err = cart.RemoveItem(itemID, quantity)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
            return
        }

        c.JSON(http.StatusOK, cart.GetCart())
    })

    // GET endpoint to get the current cart
    router.GET("/cart", func(c *gin.Context) {
        c.JSON(http.StatusOK, cart.GetCart())
    })

    // Start the server
    router.Run(":8080")
}
