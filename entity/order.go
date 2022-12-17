package entity

import (
	"gorm.io/gorm"
	"os/user"
)

type Order struct {
	gorm.Model
	UserID       uint
	User         user.User     `gorm:"foreignKey:ID;references:UserID" json:"-"`
	OrderedItems []OrderedItem `gorm:"foreignKey:OrderID"`
	TotalPrice   float32
	IsCanceled   bool
}

// NewOrder 实例化订单
func NewOrder(uid uint, items []OrderedItem) *Order {
	var totalPrice float32 = 0.0
	for _, item := range items {
		totalPrice += item.Product.Price
	}
	return &Order{
		UserID:       uid,
		OrderedItems: items,
		TotalPrice:   totalPrice,
		IsCanceled:   false,
	}
}

// BeforeCreate 创建之前，查找购物车并删除
func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {

	var cart Cart
	if err := tx.Where("UserID = ?", order.UserID).First(&cart).Error; err != nil {
		return err
	}
	if err := tx.Where("CartID = ?", cart.ID).Unscoped().Delete(&Item{}).Error; err != nil {
		return err
	}

	if err := tx.Unscoped().Delete(&cart).Error; err != nil {
		return err
	}
	return nil
}

// BeforeUpdate 如果订单被取消，金额将返回产品库存
func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {

	if order.IsCanceled {
		var orderedItems []OrderedItem
		if err := tx.Where("OrderID = ?", order.ID).Find(&orderedItems).Error; err != nil {
			return err
		}
		for _, item := range orderedItems {
			var currentProduct Product
			if err := tx.Where("ID = ?", item.ProductID).First(&currentProduct).Error; err != nil {
				return err
			}
			newStockCount := currentProduct.StockCount + item.Count
			if err := tx.Model(&currentProduct).Update(
				"StockCount", newStockCount).Error; err != nil {
				return err
			}
			if err := tx.Model(&item).Update(
				"IsCanceled", true).Error; err != nil {
				return err
			}
		}
	}
	return

}
