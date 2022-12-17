package controller

import (
	"net/http"
	"os"
	"shopping/config"
	jwtHelper "shopping/config"
	"shopping/entity"
	"shopping/enum"
	"shopping/service"
	"shopping/vo"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
	appConfig   *config.Configuration
}

// NewUserController 实例化
func NewUserController(service *service.UserService, appConfig *config.Configuration) *UserController {
	return &UserController{
		userService: service,
		appConfig:   appConfig,
	}
}

// InsertUser godoc
// @Summary 根据给定的用户名和密码创建用户
func (c *UserController) InsertUser(g *gin.Context) {
	var req vo.CreateUserRequest
	//取值
	if err := g.ShouldBind(req); err != nil {
		enum.HandleError(g, err)
		return
	}
	user := entity.NewUser(req.Username, req.Password, req.Password2)
	err := c.userService.InsertUser(user)
	if err != nil {
		enum.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, entity.GetTrueCommonResult(enum.TRUE, req.Username))
}

// Login godoc
// @Summary 根据用户名和密码登录
func (c *UserController) Login(g *gin.Context) {
	//获取表单数据
	var req vo.LoginRequest
	if err := g.ShouldBind(req); err != nil {
		enum.HandleError(g, err)
		return
	}
	//然后校验
	user, err := c.userService.FindUserByUserNameAndPassword(req.Username, req.Password)
	if err != nil {
		enum.HandleError(g, err)
		return
	}
	token := config.VerifyToken(user.Token, c.appConfig.SecretKey)
	if token == nil {
		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(user.ID), 10),
				"username": user.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp": time.Now().Add(
					24 *
						time.Hour).Unix(),
				"isAdmin": user.IsAdmin,
			})
		token := jwtHelper.GenerateToken(jwtClaims, c.appConfig.SecretKey)
		user.Token = token
		err = c.userService.UpdateByUser(&user)
		if err != nil {
			enum.HandleError(g, err)
			return
		}
	}
	g.JSON(http.StatusOK, entity.GetTrueCommonResult(enum.TRUE, user))
}

// VerifyToken 验证token
func (c *UserController) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	decodedClaims := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)
	g.JSON(http.StatusOK, decodedClaims)
}
