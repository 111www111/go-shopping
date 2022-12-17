package service

import (
	"shopping/dao"
	"shopping/entity"
	"shopping/enum"
	"shopping/utils"
	"time"
)

var day14ToHours float64 = 336

type OrderService struct {
	orderDao       dao.OrderDao
	orderedItemDao dao.OrderedItemDao
	productDao     dao.ProductDao
	cartDao        dao.CartDao
	itemDao        dao.ItemDao
}

// NewOrderService 实例化
func NewOrderService(
	orderDao dao.OrderDao,
	orderedItemDao dao.OrderedItemDao,
	productDao dao.ProductDao,
	cartDao dao.CartDao,
	itemDao dao.ItemDao,
) *OrderService {
	orderDao.Migration()
	orderedItemDao.Migration()
	return &OrderService{
		orderDao:       orderDao,
		orderedItemDao: orderedItemDao,
		productDao:     productDao,
		cartDao:        cartDao,
		itemDao:        itemDao,
	}

}

// CompleteOrder 完成订单
func (orderService *OrderService) CompleteOrder(userId uint) error {
	currentCart, err := orderService.cartDao.FindOrInsertOneByUserId(userId)
	if err != nil {
		return err
	}
	cartItems, err := orderService.itemDao.FindAllListByCartId(currentCart.UserId)
	if err != nil {
		return err
	}
	if len(cartItems) == 0 {
		return enum.ErrEmptyCartFound
	}
	orderedItems := make([]entity.OrderedItem, 0)
	for _, item := range cartItems {
		orderedItems = append(orderedItems, *entity.NewOrderedItem(item.Count, item.ProductID))
	}
	err = orderService.orderDao.Insert(entity.NewOrder(userId, orderedItems))
	return err
}

// CancelOrder 取消订单
func (orderService *OrderService) CancelOrder(uid, oid uint) error {
	currentOrder, err := orderService.orderDao.FindOneByOrderID(oid)
	if err != nil {
		return err
	}
	if currentOrder.UserID != uid {
		return enum.ErrInvalidOrderID
	}
	if currentOrder.CreatedAt.Sub(time.Now()).Hours() > day14ToHours {
		return enum.ErrCancelDurationPassed
	}
	currentOrder.IsCanceled = true
	err = orderService.orderDao.Update(*currentOrder)

	return err
}

// GetAll 获得订单
func (orderService *OrderService) GetAll(page *utils.Pages, uid uint) *utils.Pages {
	orders, count := orderService.orderDao.FindPageAllList(page.Page, page.PageSize, uid)
	page.Items = orders
	page.TotalCount = count
	return page
}
