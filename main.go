package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", controllers.GetAllProduts)
	router.GET("/products/:merchantID", controllers.GetAllProduts)
	router.POST("/products", controllers.CreateProduct)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.DeleteProduct)

	router.Run(":9090")
}
