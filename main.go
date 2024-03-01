package main

import (
	"github.com/gin-gonic/gin"

)

func main() {
	router := gin.Default()

	router.GET("/products", controllers.GetAllProducts)
	router.GET("/products/:merchantID", GetProductsByMerchant)
	router.POST("/products", CreateProduct)
	router.PUT("/products/:id", UpdatedProduct)
	router.DELETE("/products/:id", DeleteProduct)

	router.Run(":9090")
}
