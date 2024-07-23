package main

import (
	"fmt"
	"gKafkaAdmin/internal/config"
	"gKafkaAdmin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化gin
	r := gin.Default()

	// 初始化app配置
	var config config.AppConfig
	appConfigInfo := config.InitConfiguration()

	// 配置请求路由
	router.SetRouter(r)
	// 配置静态资源目录
	router.SetStaticInfo(r, *appConfigInfo)

	// 启动gin
	r.Run(fmt.Sprintf(":%d", appConfigInfo.App.Server.Port))
}
