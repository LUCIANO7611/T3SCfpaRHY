// 代码生成时间: 2025-09-15 12:47:37
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// ShoppingCart represents a simple shopping cart
type ShoppingCart struct {
    Items map[string]int
}

// AddItem adds an item to the shopping cart
func (cart *ShoppingCart) AddItem(item string, quantity int) error {
    if quantity <= 0 {
        return fmt.Errorf("invalid quantity: %d", quantity)
    }
    cart.Items[item] = cart.Items[item] + quantity
    return nil
}

// RemoveItem removes an item from the shopping cart
func (cart *ShoppingCart) RemoveItem(item string) error {
    if _, exists := cart.Items[item]; !exists {
        return fmt.Errorf("item '%s' not found in cart", item)
    }
    delete(cart.Items, item)
    return nil
}

// GetCart returns the current state of the shopping cart
func (cart *ShoppingCart) GetCart() map[string]int {
    return cart.Items
}

func main() {
    r := gin.Default()

    // Use middleware to log requests
    r.Use(gin.Logger())

    // Use middleware to recover from any panics and to set a response
    r.Use(gin.Recovery())

    cart := ShoppingCart{Items: make(map[string]int)}

    // Endpoint to add an item to the cart
    r.POST("/add", func(c *gin.Context) {
        var item struct {
            Item    string `json:"item"`
            Quantity int    `json:"quantity"`
        }
        if err := c.BindJSON(&item); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }
        if err := cart.AddItem(item.Item, item.Quantity); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(200, gin.H{
            "message": "Item added successfully",
        })
    })

    // Endpoint to remove an item from the cart
    r.POST("/remove", func(c *gin.Context) {
        var item struct {
            Item string `json:"item"`
        }
        if err := c.BindJSON(&item); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }
        if err := cart.RemoveItem(item.Item); err != nil {
            c.JSON(400, gin.H{
                "error": err.Error(),
            })
            return
        }
        c.JSON(200, gin.H{
            "message": "Item removed successfully",
        })
    })

    // Endpoint to get the current state of the cart
    r.GET("/cart", func(c *gin.Context) {
        c.JSON(200, cart.GetCart())
    })

    // Start the server
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
