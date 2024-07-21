package apiserver

import (
	"gKafkaAdmin/internal/apiserver/service"
	"gKafkaAdmin/internal/config"

	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	apiV1 := engine.Group("v1")
	{
		apiV1.GET("/test", service.TestFunc)
	}
}

func SetStaticInfo(engine *gin.Engine, appConfigInfo config.AppConfig) {
	engine.Static("/static", appConfigInfo.App.Server.StaticPath)
}
