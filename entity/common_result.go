package entity

import "errors"

// CommonResult 响应结构体
type CommonResult struct {
	Success bool        `json:"succeed"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetTrueCommonResult(mag string, data interface{}) *CommonResult {
	return &CommonResult{
		Success: true,
		Message: mag,
		Data:    data,
	}
}

func GetFalseCommonResult(mag string, data interface{}) *CommonResult {
	return &CommonResult{
		Success: false,
		Message: mag,
		Data:    data,
	}
}

// ErrInvalidBody 自定义错误
var (
	ErrInvalidBody = errors.New("请检查你的请求体")
)
