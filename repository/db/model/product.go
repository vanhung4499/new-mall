package model

import (
	"gorm.io/gorm"
	"new-mall/repository/cache"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryID    uint   `gorm:"not null"`
	Title         string
	Info          string `gorm:"size:1000"`
	ImagePath     string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.RedisContext, cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	// Increase view clicks
	cache.RedisClient.Incr(cache.RedisContext, cache.ProductViewKey(product.ID))
	// Increase ranking clicks
	cache.RedisClient.ZIncrBy(cache.RedisContext, cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
