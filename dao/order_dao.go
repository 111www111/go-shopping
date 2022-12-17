package dao

import (
	"gorm.io/gorm"
	"log"
	"shopping/entity"
)

type OrderDao struct {
	db *gorm.DB
}

// NewOrderDao 实例化
func NewOrderDao(db *gorm.DB) *OrderDao {
	return &OrderDao{
		db: db,
	}
}

// Migration 创建表
func (orderDao *OrderDao) Migration() {
	err := orderDao.db.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Print(err)
	}
}

// FindOneByOrderID 根据订单id查找
func (orderDao *OrderDao) FindOneByOrderID(oid uint) (*entity.Order, error) {
	var currentOrder *entity.Order
	if err := orderDao.db.
		Where("IsCanceled = ?", false).
		Where("ID", oid).First(&currentOrder).
		Error; err != nil {
		return nil, err
	}
	return currentOrder, nil

}

// Update 更新订单
func (orderDao *OrderDao) Update(newOrder entity.Order) error {
	result := orderDao.db.Save(&newOrder)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Insert 创建订单
func (orderDao *OrderDao) Insert(ci *entity.Order) error {
	result := orderDao.db.Create(ci)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindPageAllList 获得所有订单
func (orderDao *OrderDao) FindPageAllList(pageIndex, pageSize int, uid uint) ([]entity.Order, int) {
	var orders []entity.Order
	var count int64
	orderDao.db.Where("IsCanceled = ?", 0).Where(
		"UserID", uid).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&orders).Count(&count)
	for i, order := range orders {
		orderDao.db.Where("OrderID = ?", order.ID).Find(&orders[i].OrderedItems)
		for j, item := range orders[i].OrderedItems {
			orderDao.db.Where("ID = ?", item.ProductID).First(&orders[i].OrderedItems[j].Product)
		}
	}
	return orders, int(count)
}
