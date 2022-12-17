package dao

import (
	"gorm.io/gorm"
	"log"
	"shopping/entity"
)

type CategoryDao struct {
	Db *gorm.DB
}

func NewCategoryDao(db *gorm.DB) *CategoryDao {
	return &CategoryDao{
		Db: db,
	}
}

// Insert 创建商品类型
func (categoryDao *CategoryDao) Insert(category *entity.Category) {
	categoryDao.Db.Create(category)
}

// InsertByList 批量创建
func (categoryDao *CategoryDao) InsertByList(category []*entity.Category) (int, error) {
	var count int64
	err := categoryDao.Db.
		Create(&category).
		Count(&count).Error
	return int(count), err
}

// FindListByName 通过名称查询商品分类
func (categoryDao *CategoryDao) FindListByName(name string) []entity.Category {
	var returnList []entity.Category
	categoryDao.Db.
		Where("name = ?", name).
		Find(&returnList)
	return returnList
}

// FindListPage 获得分页商品分类
func (categoryDao *CategoryDao) FindListPage(pageIndex, pageSize int) (int, []entity.Category) {
	var returnList []entity.Category
	var count int64
	categoryDao.Db.
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Find(&returnList).
		Count(&count)
	return int(count), returnList
}

// Migration 维护表结构
func (categoryDao *CategoryDao) Migration() {
	err := categoryDao.Db.AutoMigrate(&entity.Category{})
	if err != nil {
		log.Print(err)
	}
}

// FindOneByName 查询根据name
func (categoryDao *CategoryDao) FindOneByName(name string) entity.Category {
	var returnCategory entity.Category
	categoryDao.Db.Where("name = ?", name).First(&returnCategory)
	return returnCategory
}
