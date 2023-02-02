package controllers

import (
	"cros-rate-limit-service/internal/dto"
	"github.com/gin-gonic/gin"
)

type CrosRiskEngineController struct {
	Base
}

func (m *CrosRiskEngineController) Decision(c *gin.Context) {
	// 处理入参
	var request dto.DecisionReq
	err := c.ShouldBindJSON(&request)
	if err != nil {
		m.Fail(c, "", "")
		return
	}

	// 根据服务名、决策流key、版本 获取到决策流

	// 根据features入参 匹配json schema 报错返回

	// 根据决策流需要的特征 从features取出 set到pipeline ctx中

	// 遍历node节点 执行决策
}
