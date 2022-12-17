package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//TODO: refactor methods
var (
	// DefaultPageSize 默认页数
	DefaultPageSize = 100
	// MaxPageSize 最大页数
	MaxPageSize = 1000
	// PageVar 查询参数名称
	PageVar = "page"
	// PageSizeVar 页数查询参数名称
	PageSizeVar = "pageSize"
)

// Pages 分页结构体
type Pages struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	PageCount  int         `json:"pageCount"`
	TotalCount int         `json:"totalCount"`
	Items      interface{} `json:"items"`
}

// New 实例化分页结构体
func New(page, pageSize, total int) *Pages {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	if pageSize > MaxPageSize {
		pageSize = MaxPageSize
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + pageSize - 1) / pageSize
		if page > pageCount {
			page = pageCount
		}

	}
	if page <= 0 {
		page = 1
	}

	return &Pages{
		Page:       page,
		PageSize:   pageSize,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// NewFromRequest 根据http请求实例化分页结构体
func NewFromRequest(req *http.Request, count int) *Pages {
	page := ParseInt(req.URL.Query().Get(PageVar), 1)
	pageSize := ParseInt(req.URL.Query().Get(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// NewFromGinRequest 根据gin请求实例化分页结构体
func NewFromGinRequest(g *gin.Context, count int) *Pages {
	page := ParseInt(g.Query(PageVar), 1)
	pageSize := ParseInt(g.Query(PageSizeVar), DefaultPageSize)
	return New(page, pageSize, count)
}

// ParseInt 类型转换
func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

// Offset offset
func (p *Pages) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// Limit limit
func (p *Pages) Limit() int {
	return p.PageSize
}
