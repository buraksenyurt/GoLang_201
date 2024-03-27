package main

import (
	"example/databaseops/models"
	"fmt"
)

func main() {
	product := models.Product{
		Title:       "Rust Programming Language",
		Description: "A simple rust book",
		Price:       34.50,
	}
	var id, title = models.InsertProduct(product)

	fmt.Printf("Inserted row (%d) - %s\n", id, title)

	product.Id = id
	product.Title = "Sustainabile Software Engineering"
	product.Description = "A really nice book"
	product.Price = 59.99
	models.UpdateProduct(product)

	models.GetProducts()

	models.GetProductById(id)
}
