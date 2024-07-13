package main

import (
	"fmt"
	"gKafkaAdmin/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var config config.AppConfig
	appConfigInfo := config.InitConfiguration()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "TEST OK")
	})

	r.Run(fmt.Sprintf(":%d", appConfigInfo.App.Server.Port))
}
