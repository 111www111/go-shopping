package entity

import "gorm.io/gorm"

//Category 商品
type Category struct {
	gorm.Model
	//唯一
	Name string `gorm:"unique"`
	Desc string
	//是否激活
	IsActive bool
}

// NewCategory 新建商品分类
func NewCategory(name, desc string) *Category {
	return &Category{
		Name: name,
		Desc: desc,
	}
}
