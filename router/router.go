package router

import (
	"cros-rate-limit-service/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(engine *gin.Engine) *gin.Engine {
	// 缺失路由
	engine.NoRoute(NotFoundHandler)

	// 健康检查
	engine.GET("/checkhealth", CheckHealth)

	// swag 文档
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 限流服务路由
	c := controllers.CrosRiskEngineController{}
	group := engine.Group("/api/v1")
	group.POST("/decision", c.Decision)
	return engine
}
