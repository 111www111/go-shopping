package service

import (
	"mime/multipart"
	"shopping/dao"
	"shopping/entity"
	"shopping/enum"
	"shopping/utils"
)

type CategoryService struct {
	categoryDao dao.CategoryDao
}

// NewCategoryService 实例化
func NewCategoryService(categoryDao dao.CategoryDao) *CategoryService {
	categoryDao.Migration()
	return &CategoryService{
		categoryDao: categoryDao,
	}
}

// Create 创建分类
func (c *CategoryService) Create(category *entity.Category) error {
	if existCity := c.categoryDao.FindOneByName(category.Name); existCity != (entity.Category{}) {
		return enum.CategoryErrExistWithName
	}
	err := c.categoryDao.Db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

// BulkCreate 批量创建分类
func (c *CategoryService) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*entity.Category, 0)
	bulkCategory, err := utils.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, entity.NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	count, err := c.categoryDao.InsertByList(categories)
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetAll 获得分页商品分类
func (c *CategoryService) GetAll(page *utils.Pages) *utils.Pages {
	count, categories := c.categoryDao.FindListPage(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page
}
