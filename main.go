package main

import (
	controller "./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", controller.GetAllProducts)
	router.GET("/products/:merchantID", controller.GetProductsByMerchant)
	router.POST("/products", controller.CreateProduct)
	router.PUT("/products/:id", controller.UpdatedProduct)
	router.DELETE("/products/:id", controller.DeleteProduct)

	router.Run(":9090")
}
