package entity

import (
	"gorm.io/gorm"
	"shopping/enum"
)

// OrderedItem 订单项结构体
type OrderedItem struct {
	gorm.Model
	Product    Product `gorm:"foreignKey:ProductID"`
	ProductID  uint
	Count      int
	OrderID    uint
	IsCanceled bool
}

// NewOrderedItem 实例化订单项
func NewOrderedItem(count int, pid uint) *OrderedItem {
	return &OrderedItem{
		Count:      count,
		ProductID:  pid,
		IsCanceled: false,
	}
}

// BeforeSave 保存之前，更新产品库存
func (orderedItem *OrderedItem) BeforeSave(tx *gorm.DB) (err error) {

	var currentProduct Product
	var currentOrderedItem OrderedItem
	if err := tx.Where("ID = ?", orderedItem.ProductID).First(&currentProduct).Error; err != nil {
		return err
	}
	reservedStockCount := 0
	if err := tx.Where("ID = ?", orderedItem.ID).First(&currentOrderedItem).Error; err == nil {
		reservedStockCount = currentOrderedItem.Count
	}
	newStockCount := currentProduct.StockCount + reservedStockCount - orderedItem.Count
	if newStockCount < 0 {
		return enum.ErrNotEnoughStock
	}
	if err := tx.Model(&currentProduct).Update("StockCount", newStockCount).Error; err != nil {
		return err
	}
	if orderedItem.Count == 0 {
		err := tx.Unscoped().Delete(currentOrderedItem).Error
		return err
	}
	return
}
