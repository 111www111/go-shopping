package service

import (
	"errors"
	"shopping/dao"
	"shopping/entity"
	"shopping/enum"
)

type CartService struct {
	cartDao    dao.CartDao
	itemDao    dao.ItemDao
	productDao dao.ProductDao
}

// NewCartService 实例化service
func NewCartService(cartDao dao.CartDao, itemDao dao.ItemDao, productDao dao.ProductDao) *CartService {
	cartDao.Migration()
	itemDao.Migration()
	return &CartService{
		cartDao:    cartDao,
		itemDao:    itemDao,
		productDao: productDao,
	}
}

// InsertItem 添加item
func (cartService *CartService) InsertItem(userID uint, sku string, count int) error {
	currentProduct, err := cartService.productDao.FindOneBySKU(sku)
	if err != nil {
		return err
	}
	currentCart, err := cartService.cartDao.FindOrInsertOneByUserId(userID)
	if err != nil {
		return err
	}
	_, err = cartService.itemDao.FindOneById(currentProduct.ID, currentCart.ID)
	if err == nil {
		return enum.ItemErrAlreadyExistInCart
	}
	if currentProduct.StockCount < count {
		return enum.ProductErrStockIsNotEnough
	}
	if count <= 0 {
		return enum.CartErrCountInvalid
	}
	err = cartService.itemDao.Insert(entity.NewCartItem(currentProduct.ID, currentCart.ID, count))
	return err
}

// UpdateItem 更新item
func (cartService *CartService) UpdateItem(userID uint, sku string, count int) error {
	currentProduct, err := cartService.productDao.FindOneBySKU(sku)
	if err != nil {
		return err
	}
	currentCart, err := cartService.cartDao.FindOrInsertOneByUserId(userID)
	if err != nil {
		return err
	}
	currentItem, err := cartService.itemDao.FindOneById(currentProduct.ID, currentCart.ID)
	if err != nil {
		return errors.New("item 不存在")
	}
	if currentProduct.StockCount+currentItem.Count < count {
		return enum.ProductErrStockIsNotEnough
	}
	currentItem.Count = count
	err = cartService.itemDao.UpdateById(*currentItem)
	return err
}

// GetCartItems 获得items
func (cartService *CartService) GetCartItems(userId uint) ([]entity.Item, error) {
	currentCart, err := cartService.cartDao.FindOrInsertOneByUserId(userId)
	if err != nil {
		return nil, err
	}
	items, err := cartService.itemDao.FindAllListByCartId(currentCart.ID)
	if err != nil {
		return nil, err
	}
	return items, nil
}
