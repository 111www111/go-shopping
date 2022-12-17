package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId uint
	User   User `gorm:"foreignKey:ID;references:UserID"`
}

func NewCart(userId uint) *Cart {
	return &Cart{
		UserId: userId,
	}
}

//AfterUpdate 切面如果计数为零，则删除商品
func (item *Item) AfterUpdate(tx *gorm.DB) (err error) {

	if item.Count <= 0 {
		return tx.Unscoped().Delete(&item).Error
	}
	return
}
