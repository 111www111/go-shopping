package enum

import (
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"net/http"
	"shopping/entity"
)

var (
	UserErrUserExistWithName   = errors.New("用户名已经存在")
	UserErrUserNotFound        = errors.New("用户名或密码错误")
	UserErrMismatchedPasswords = errors.New("密码不匹配")
	UserErrInvalidUsername     = errors.New("无效用户名")
	UserErrInvalidPassword     = errors.New("无效密码")
)

var (
	CategoryErrExistWithName = errors.New("商品分类已经存在")
)

var (
	ProductErrNotFound         = errors.New("商品没有找到")
	ProductErrStockIsNotEnough = errors.New("商品库存不足")
)

var (
	ItemErrAlreadyExistInCart = errors.New("商品已经存在")
	CartErrCountInvalid       = errors.New("数量不能是负值")
)

var (
	ErrEmptyCartFound       = errors.New("购物车是空的")
	ErrInvalidOrderID       = errors.New("无效订单")
	ErrCancelDurationPassed = errors.New("已通过取消持续时间")
	ErrNotEnoughStock       = errors.New("没有足够库存")
)

// HandleError 错误处理
func HandleError(g *gin.Context, err error) {
	log.Errorf(g, "error: %v", err)
	g.JSON(
		http.StatusBadRequest, entity.GetFalseCommonResult(Err500Msg, nil))
	g.Abort()
	return
}
