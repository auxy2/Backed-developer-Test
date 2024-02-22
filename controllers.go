package main

import (
	// "errors"
	"github.com/gin-gonic/gin"
	"time"
)

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

var products = []Product{}

func getAllProduts(c *gin.Context) {
	c.IndentedJSON(200, products)
}

func getProductsByMerchant(c *gin.Context) {
	merchantID := c.Param("merchantID")

	merchantProducts := []Product{}
	for _, p := range products {
		if p.ID == merchantID {
			merchantProducts = append(merchantProducts, p)
		}
	}

	c.JSON(200, merchantProducts)
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	newProduct.CreatedAt = time.Now()
	products = append(products, newProduct)

	c.JSON(201, newProduct)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")

	var updatedProduct Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(400, gin.H{"message": "Page not Found"})
		return
	}

	var found bool
	for i, p := range products {
		if p.ID == id {
			updatedProduct.CreatedAt = p.CreatedAt
			products[i] = updatedProduct
			found = true
			break
		}
	}

	if !found {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}

	c.JSON(200, updatedProduct)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")

	var found bool
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		c.JSON(404, gin.H{"message": "Product not found"})
		return
	}

	c.JSON(204, nil)
}
