package main

import (
	"example/databaseops/models"
)

func main() {
	product := models.Product{
		Title:       "Rust Programming Language",
		Description: "A simple rust book",
		Price:       34.50,
	}
	models.InsertProduct(product)
}
