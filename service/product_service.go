package service

import (
	"shopping/dao"
	"shopping/entity"
	"shopping/utils"
)

type ProductService struct {
	productDao dao.ProductDao
}

// NewService 实例化
func NewService(productDao dao.ProductDao) *ProductService {
	productDao.Migration()
	return &ProductService{
		productDao: productDao,
	}

}

// FindList 获得所有商品分页
func (c *ProductService) FindList(page *utils.Pages) *utils.Pages {
	products, count := c.productDao.FindList(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}

// Insert 创建商品
func (c *ProductService) Insert(name string, desc string, count int, price float32, cid uint) error {
	newProduct := entity.NewProduct(name, desc, count, price, cid)
	err := c.productDao.Insert(newProduct)
	return err
}

// DeleteBySKU 删除商品
func (c *ProductService) DeleteBySKU(sku string) error {
	err := c.productDao.DeleteBySKU(sku)
	return err
}

// UpdateProduct 更新商品
func (c *ProductService) UpdateProduct(product *entity.Product) error {
	err := c.productDao.UpdateById(*product)
	return err
}

// FindPageListByCondition 查找商品
func (c *ProductService) FindPageListByCondition(text string, page *utils.Pages) *utils.Pages {
	products, count := c.productDao.FindPageListByCondition(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}
