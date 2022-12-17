package entity

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Product   Product `gorm:"foreignKey:ProductID"`
	ProductID uint
	Count     int
	CartID    uint
	Cart      Cart `gorm:"foreignKey:CartID" json:"-"`
}

// NewCartItem 创建Item
func NewCartItem(productId uint, cartId uint, count int) *Item {
	return &Item{
		ProductID: productId,
		Count:     count,
		CartID:    cartId,
	}
}
