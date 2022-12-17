package entity

import (
	"gorm.io/gorm"
	"shopping/utils"
)

// User 用户模型
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30)"`
	Password  string `gorm:"type:varchar(100)"`
	Password2 string `gorm:"-"`
	Salt      string `gorm:"type:varchar(100)"`
	Token     string `gorm:"type:varchar(500)"`
	IsDeleted bool
	IsAdmin   bool
}

// NewUser 构造方法
func NewUser(username, password, password2 string) *User {
	return &User{
		Username:  username,
		Password:  password,
		Password2: password2,
		IsDeleted: false,
		IsAdmin:   false,
	}
}

// BeforeSave aop方法,如果用户回调前没有加密，则加密密码
func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Salt == "" {
		//创建一个随机字符串
		salt := utils.CreateSalt()
		//创建密码加密
		password, err := utils.HashPassword(u.Password + salt)
		if err != nil {
			return err
		}
		u.Password = password
		u.Salt = salt
	}
	return nil
}
