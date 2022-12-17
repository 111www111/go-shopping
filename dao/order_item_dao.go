package dao

import (
	"gorm.io/gorm"
	"log"
	"shopping/entity"
)

type OrderedItemDao struct {
	db *gorm.DB
}

// NewOrderedItemDao 实例化
func NewOrderedItemDao(db *gorm.DB) *OrderedItemDao {
	return &OrderedItemDao{
		db: db,
	}
}

// Migration 创建表
func (orderedItemDao *OrderedItemDao) Migration() {
	err := orderedItemDao.db.AutoMigrate(&entity.OrderedItem{})
	if err != nil {
		log.Print(err)
	}
}

// Update 更新
func (orderedItemDao *OrderedItemDao) Update(item entity.OrderedItem) error {
	result := orderedItemDao.db.Save(&item)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Insert 创建订单item
func (orderedItemDao *OrderedItemDao) Insert(ci *entity.OrderedItem) error {
	result := orderedItemDao.db.Create(ci)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
