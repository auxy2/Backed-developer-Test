package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", controllers.getAllProducts)
	router.GET("/products/:merchantID", getProductsByMerchant)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updatedProduct)
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":9090")
}
