package main

import (
	"fmt"
	"gKafkaAdmin/internal/config"
	"gKafkaAdmin/internal/zlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var config config.AppConfig
	appConfigInfo := config.InitConfiguration()

	r.GET("/test", func(ctx *gin.Context) {
		zlog.Info("test.OK")
		ctx.String(http.StatusOK, "TEST OK")
	})

	r.Run(fmt.Sprintf(":%d", appConfigInfo.App.Server.Port))
}
