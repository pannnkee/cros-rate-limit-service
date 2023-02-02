package controllers

import (
	"cros-rate-limit-service/internal/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	// Resp 响应
	Resp struct {
		Code int         `json:"code"` // 应答码
		Msg  string      `json:"msg"`  // 应答信息
		Data interface{} `json:"data"` // 应答数据
	}

	// RespPage 分页相应
	RespPage struct {
		Resp
		Total    int64 `json:"total"`     // 分页总数
		Page     int64 `json:"page"`      // 当前页
		PageSize int64 `json:"page_size"` // 每页数量
	}

	Base struct{}
)

func (m *Base) Success(c *gin.Context, message string, data interface{}) {
	if message == "" {
		message = errcode.ResponseMap[errcode.Success]
	}
	r := Resp{
		Code: errcode.Success,
		Msg:  message,
		Data: data,
	}
	c.JSON(http.StatusOK, r)
}

func (m *Base) Fail(c *gin.Context, message string, data interface{}) {
	if message == "" {
		message = errcode.ResponseMap[errcode.Failed]
	}
	r := Resp{
		Code: errcode.Failed,
		Msg:  message,
		Data: data,
	}
	c.JSON(http.StatusOK, r)
}
