package dao

import (
	"gorm.io/gorm"
	"log"
	"shopping/entity"
)

type CartDao struct {
	db *gorm.DB
}

func NewCartDao(db *gorm.DB) *CartDao {
	return &CartDao{db: db}
}

// Migration 创建表
func (cartDao *CartDao) Migration() {
	err := cartDao.db.AutoMigrate(&entity.Cart{})
	if err != nil {
		log.Print(err)
	}
}

//UpdateById 更新
func (cartDao *CartDao) UpdateById(cart *entity.Cart) error {
	if err := cartDao.db.Save(&cart).Error; err != nil {
		return err
	}
	return nil
}

// FindOrInsertOneByUserId 根据用户id查找或创建购物车
func (cartDao *CartDao) FindOrInsertOneByUserId(userId uint) (*entity.Cart, error) {
	var cart *entity.Cart
	err := cartDao.db.Where(entity.Cart{UserId: userId}).Attrs(entity.NewCart(userId)).FirstOrCreate(&cart).Error
	if err != nil {
		return nil, err
	}
	return cart, nil
}

// FindOneByUserId 根据用户id查找购物车
func (cartDao *CartDao) FindOneByUserId(userId uint) (*entity.Cart, error) {
	var cart *entity.Cart
	if err := cartDao.db.Where(entity.Cart{UserId: userId}).First(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}
