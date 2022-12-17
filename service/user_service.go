package service

import (
	"github.com/gin-gonic/gin"
	"shopping/dao"
	"shopping/entity"
	"shopping/enum"
	"shopping/utils"
)

type UserService struct {
	userDao dao.UserDao
}

// NewUserService 实例化service
func NewUserService(userDao dao.UserDao) *UserService {
	userDao.Migration()
	return &UserService{
		userDao: userDao,
	}
}

//InsertUser 创建用户
func (userService *UserService) InsertUser(user *entity.User) error {
	//用户两次输入密码不一致
	if user.Password != user.Password2 {
		return enum.UserErrMismatchedPasswords
	}
	//查看用户是否被重命名
	userByDb, _ := userService.userDao.FindOneByUserName(user.Username)
	if userByDb != (entity.User{}) {
		//这里说明userByDb有值
		return enum.UserErrUserExistWithName
	}
	//验证用户的密码和用户名是否符合正则
	if utils.ValidateUserName(user.Username) {
		return enum.UserErrInvalidUsername
	}
	if utils.ValidatePassword(user.Password) {
		return enum.UserErrInvalidPassword
	}
	//创建用户
	err := userService.userDao.Insert(user)
	return err
}

//FindUserByUserNameAndPassword 查询用户
func (userService *UserService) FindUserByUserNameAndPassword(username, password string) (entity.User, error) {
	user, err := userService.userDao.FindOneByUserName(username)
	if err != nil {
		return entity.User{}, err
	}
	isIdentical := utils.CheckPasswordHash(password, user.Salt)
	if !isIdentical {
		return entity.User{}, enum.UserErrUserNotFound
	}
	return user, nil
}

//UpdateByUser 更新用户信息
func (userService *UserService) UpdateByUser(user *entity.User) error {
	return userService.userDao.Update(user)
}

var userIdText = "userId"

// GetUserId 从context获得用户id
func GetUserId(g *gin.Context) uint {
	return uint(utils.ParseInt(g.GetString(userIdText), -1))
}
