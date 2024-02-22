package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", getAllProduts)
	router.GET("/products/:merchantID", getProductsByMerchant)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	router.Run(":9090")
}
