package main

import (
	"fmt"
	"gKafkaAdmin/internal/apiserver"
	"gKafkaAdmin/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var config config.AppConfig
	appConfigInfo := config.InitConfiguration()

	apiserver.SetRouter(r)

	r.Run(fmt.Sprintf(":%d", appConfigInfo.App.Server.Port))
}
