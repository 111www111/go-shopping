package dao

import (
	"gorm.io/gorm"
	"log"
	"shopping/entity"
	"shopping/enum"
)

type ProductDao struct {
	db *gorm.DB
}

func NewProductDao(db *gorm.DB) *ProductDao {
	return &ProductDao{
		db: db,
	}
}

// Migration 生成表
func (r *ProductDao) Migration() {
	err := r.db.AutoMigrate(&entity.Product{})
	if err != nil {
		log.Print(err)
	}
}

// UpdateById Update 更新
func (r *ProductDao) UpdateById(updateProduct entity.Product) error {
	savedProduct, err := r.FindOneBySKU(updateProduct.SKU)
	if err != nil {
		return err
	}
	err = r.db.Model(&savedProduct).Updates(updateProduct).Error
	return err
}

// FindPageListByCondition 搜索返回分页结果
func (r *ProductDao) FindPageListByCondition(str string, pageIndex, pageSize int) ([]entity.Product, int) {
	var products []entity.Product
	convertedStr := "%" + str + "%"
	var count int64
	r.db.Where("IsDeleted = ?", false).Where(
		"Name LIKE ? OR SKU Like ?", convertedStr,
		convertedStr).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

// FindOneBySKU 根据sku查找
func (r *ProductDao) FindOneBySKU(sku string) (*entity.Product, error) {
	var product *entity.Product
	err := r.db.Where("IsDeleted = ?", 0).Where(entity.Product{SKU: sku}).First(&product).Error
	if err != nil {
		return nil, enum.ProductErrNotFound
	}
	return product, nil
}

// Insert 创建
func (r *ProductDao) Insert(p *entity.Product) error {
	result := r.db.Create(p)

	return result.Error
}

// FindList 查询所有商品
func (r *ProductDao) FindList(pageIndex, pageSize int) ([]entity.Product, int) {
	var products []entity.Product
	var count int64

	r.db.Where("IsDeleted = ?", 0).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&products).Count(&count)

	return products, int(count)
}

// DeleteBySKU 根据sku删除
func (r *ProductDao) DeleteBySKU(sku string) error {
	currentProduct, err := r.FindOneBySKU(sku)
	if err != nil {
		return err
	}
	currentProduct.IsDeleted = true

	err = r.db.Save(currentProduct).Error
	return err
}
