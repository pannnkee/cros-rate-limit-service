package router

import (
	"cros-rate-limit-service/configs"
	"cros-rate-limit-service/controllers"
	"cros-rate-limit-service/internal/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CheckHealth
// @Description 探针路由
// @Summary 决策系统探针接口
// @Router /checkhealth [GET]
// @Success 200 {object} controllers.Resp
func CheckHealth(c *gin.Context) {
	time.Sleep(time.Second)
	r := controllers.Resp{
		Code: errcode.Success,
		Msg:  fmt.Sprintf("%s state health...", configs.ServiceName),
		Data: []int{},
	}
	c.JSON(http.StatusOK, r)
}
