package main

// import (
// 	// "context"
// 	"errors"

// 	"github.com/gin-gonic/gin"
// 	// "net/http"
// )

// type product struct {
// 	Id          string `json:"id"`
// 	Product     string `json:"product"`
// 	Description string `json:"description"`
// 	Price       int    `json:"price"`
// }

// var products = []product{
// 	{Id: "1", Product: "freezer", Description: "One door standing freezer", Price: 200456},
// 	{Id: "2", Product: "table", Description: "a wooden table well fuenished", Price: 10676},
// 	{Id: "3", Product: "3D light", Description: "RGB light lamb for night", Price: 9893},
// 	{Id: "4", Product: "microwave", Description: "digital microwave with parfect heat condition", Price: 40589},
// }

// func getAllProduts(context *gin.Context) {
// 	context.IndentedJSON(200, products)
// }

// func addProducts(context *gin.Context) {
// 	var newProducts product

// 	if err := context.BindJSON(&newProducts); err != nil {
// 		return
// 	}

// 	products = append(products, newProducts)
// 	context.IndentedJSON(201, newProducts)
// }

// func getProduct(context *gin.Context) {
// 	id := context.Param("id")

// 	product, err := getProdutsById(id)

// 	if err != nil {
// 		context.IndentedJSON(404, gin.H{"message": "Product Not Found"})
// 		return
// 	}

// 	context.IndentedJSON(200, product)
// }

// func getProdutsById(id string) (*product, error) {
// 	for i, p := range products {
// 		if p.Id == id {
// 			return &products[i], nil
// 		}
// 	}

// 	return nil, errors.New("product not found")
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/products", getAllProduts)
// 	router.POST("/addProducts", addProducts)
// 	router.GET("/getProduct/:id", getProduct)

// 	router.Run("localhost:9090")
// }
