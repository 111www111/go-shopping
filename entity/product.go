package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name       string
	SKU        string
	Desc       string
	StockCount int
	Price      float32
	CategoryID uint     // 分类id
	Category   Category `json:"-"` // 分类
	IsDeleted  bool
}

// NewProduct 商品结构体实例
func NewProduct(name string, desc string, stockCount int, price float32, cid uint) *Product {
	return &Product{
		Name:       name,
		Desc:       desc,
		StockCount: stockCount,
		Price:      price,
		CategoryID: cid,
		IsDeleted:  false,
	}
}

// BeforeSave 保存商品之前生成商品sku
func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	p.SKU = uuid.New().String()
	return
}
