package trace

import (
	"context"
	"cros-rate-limit-service/configs"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Trace 定义trace结构体
type Trace struct {
	Trx string `json:"trx"`
}

// GetTraceCtx 根据gin的context获取context，使log trace更加通用
func GetTraceCtx(c *gin.Context) context.Context {
	get := c.MustGet(configs.TraceCtx)
	fmt.Println(get)
	a := get.(context.Context)
	fmt.Println(a)
	return c.MustGet(configs.TraceCtx).(context.Context)
}
