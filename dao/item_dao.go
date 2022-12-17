package dao

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"shopping/entity"
)

type ItemDao struct {
	db *gorm.DB
}

func NewItemDao(db *gorm.DB) *ItemDao {
	return &ItemDao{db: db}
}

// Migration 生成item表
func (itemDao *ItemDao) Migration() {
	err := itemDao.db.AutoMigrate(&entity.Item{})
	if err != nil {
		log.Print(err)
	}
}

// UpdateById 更新item
func (itemDao *ItemDao) UpdateById(item entity.Item) error {
	result := itemDao.db.Save(&item)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindOneById 根据商品id和购物车id查找item
func (itemDao *ItemDao) FindOneById(pid uint, cid uint) (*entity.Item, error) {
	var item *entity.Item

	err := itemDao.db.Where(&entity.Item{ProductID: pid, CartID: cid}).First(&item).Error
	if err != nil {
		return nil, errors.New("cart item not found")
	}
	return item, nil
}

//Insert 创建item
func (itemDao *ItemDao) Insert(item *entity.Item) error {
	result := itemDao.db.Create(item)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindAllListByCartId 返回购物车中所有item
func (itemDao *ItemDao) FindAllListByCartId(cartId uint) ([]entity.Item, error) {
	var cartItems []entity.Item
	err := itemDao.db.Where(&entity.Item{CartID: cartId}).Find(&cartItems).Error
	if err != nil {
		return nil, err
	}
	for i, item := range cartItems {
		err := itemDao.db.Model(item).Association("Product").Find(&cartItems[i].Product)
		if err != nil {
			return nil, err
		}
	}
	return cartItems, nil
}
