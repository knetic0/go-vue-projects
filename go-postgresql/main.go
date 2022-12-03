package main

import (
	"fmt"

	"github.com/knetic0/web-develop-examples/models"
)

func main() {
	product := models.Product{
		Title:       "Go Programming Language Book",
		Description: "It's a good book",
		Price:       36.78,
	}
	fmt.Println(product)
	//models.InsertProduct(product)
	models.GetProductsByID(2)
	models.GetProducts()
	models.DeleteProductsByID(1)
}
