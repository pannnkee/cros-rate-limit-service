package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundHandler(c *gin.Context) {
	r := struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}{
		Code: http.StatusNotFound,
		Msg:  "地址不存在",
		Data: []int{},
	}
	c.JSON(http.StatusNotFound, r)
}
