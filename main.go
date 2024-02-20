package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type product struct {
	Id          string `json:"id"`
	Product     string `json:"product"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

var products = []product{
	{Id: "1", Product: "freezer", Description: "One door standing freezer", Price: 200456},
	{Id: "2", Product: "table", Description: "a wooden table well fuenished", Price: 10676},
	{Id: "3", Product: "3D light", Description: "RGB light lamb for night", Price: 9893},
	{Id: "4", Product: "microwave", Description: "digital microwave with parfect heat condition", Price: 40589},
}

func main (){
	router := gin.Default()
	router.GET("/products")
	router.Run("localhost:9090")
}
fmt.Printf("you are welcome")
