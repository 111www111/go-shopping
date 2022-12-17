package dao

import (
	"gorm.io/gorm"
	"log"
	"shopping/entity"
)

// UserDao 结构体
type UserDao struct {
	db *gorm.DB
}

// NewUserDao 实例化
func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

// Insert 添加用户
func (dao *UserDao) Insert(user *entity.User) error {
	tx := dao.db.Create(user)
	return tx.Error
}

// FindOneByUserName 根据用户名查询记录
func (dao *UserDao) FindOneByUserName(username string) (entity.User, error) {
	var user entity.User
	err := dao.db.Where("username = ?", username).Where("IsDeleted = ?", 0).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// Update 根据实体类更新
func (dao *UserDao) Update(user *entity.User) error {
	return dao.db.Save(&user).Error
}

//Migration 维护表结构
func (dao *UserDao) Migration() {
	err := dao.db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Print(err)
	}
}
