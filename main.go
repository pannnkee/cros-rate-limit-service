package main

import (
	CrosRateLimitService "cros-rate-limit-service/cmd/cros-rate-limit-service"
	"cros-rate-limit-service/serverman"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 加载配置

	// 项目启动
	serverman.RegisterServer(CrosRateLimitService.New())
	if err := serverman.Start(); err != nil {
		fmt.Println(err)
	}

}
