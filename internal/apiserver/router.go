package apiserver

import (
	testservice "gKafkaAdmin/internal/apiserver/service"
	"gKafkaAdmin/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	apiV1 := engine.Group("v1")
	{
		apiV1.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, testservice.TestFunc(ctx))
		})
		apiV1.GET("/test2", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, testservice.TestFunc2(ctx))
		})
	}
}

func SetStaticInfo(engine *gin.Engine, appConfigInfo config.AppConfig) {
	engine.Static("/static", appConfigInfo.App.Server.StaticPath)
}
