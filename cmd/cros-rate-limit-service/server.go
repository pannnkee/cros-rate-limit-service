package CrosRateLimitService

import (
	"context"
	"cros-rate-limit-service/internal/log"
	"cros-rate-limit-service/middleware"
	"cros-rate-limit-service/router"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// CrosRiskEngineService 反欺诈服务
type CrosRiskEngineService struct {
	server *http.Server
}

func New() *CrosRiskEngineService {
	return &CrosRiskEngineService{}
}

// Serve 启动服务
func (m *CrosRiskEngineService) Serve() (err error) {
	engine := gin.New()
	logger := log.NewLogger(zap.InfoLevel)
	logger = logger.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1))

	// 捕获异常
	engine.Use(middleware.GinLogger(logger), middleware.GinRecovery(logger, true))

	// 跨域
	engine.Use(middleware.Cors())

	// 设置路由
	router.InitRouter(engine)

	m.server = &http.Server{
		Addr:         ":8090",
		Handler:      engine,
		ReadTimeout:  time.Duration(60) * time.Second,
		WriteTimeout: time.Duration(60) * time.Second,
	}

	if err = m.server.ListenAndServe(); err != nil {
		return err
	}
	return
}

// Stop 退出
func (m *CrosRiskEngineService) Stop() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return m.server.Shutdown(ctx)
}
