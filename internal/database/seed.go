package database

import (
	"github.com/go-faker/faker/v4"
	"log"
	"new-mall/internal/global"
	"new-mall/internal/models"
)

func Seed() {

	// Create mock categories
	var categories []models.Category
	for i := 0; i < 5; i++ {
		var category models.Category
		err := faker.FakeData(&category)
		if err != nil {
			log.Fatal(err)
		}
		global.DB.Create(&category)
		categories = append(categories, category)
	}

	// Create mock products
	var products []models.Product
	for i := 0; i < 10; i++ {
		var product models.Product
		err := faker.FakeData(&product)
		if err != nil {
			log.Fatal(err)
		}
		global.DB.Create(&product)
		products = append(products, product)
	}
}
