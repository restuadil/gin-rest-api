package services

import "gin-app/models"

var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 1500.00},
	{ID: 2, Name: "Smartphone", Price: 800.00},
}

func GetAllProducts() []models.Product {
	return products
}
