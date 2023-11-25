package database

import (
	"new-mall/internal/global"
	"new-mall/internal/models"
)

func Migrate() {
	global.DB.AutoMigrate(
		&models.User{},
		&models.Favorite{},
		&models.Order{},
		&models.OrderItem{},
		&models.Address{},
		&models.Cart{},
		&models.CartItem{},
		&models.Category{},
		&models.Carousel{},
		&models.Notice{},
		&models.Product{},
		&models.Image{},
	)
}
