package router

import (
	"gKafkaAdmin/internal/config"
	kafkaController "gKafkaAdmin/internal/controller/kafka"
	testController "gKafkaAdmin/internal/controller/test"

	"github.com/gin-gonic/gin"
)

func SetRouter(engine *gin.Engine) {
	apiV1 := engine.Group("v1")
	{
		apiV1.GET("/test", testController.TestFunc)

		apiV1.GET("/kafka/topic/getAllTopic", kafkaController.GetTopicList)
		apiV1.GET("/kafka/topic/createTopic", kafkaController.GetTopicList)
		apiV1.GET("/kafka/topic/deleteTopic", kafkaController.GetTopicList)
	}
}

func SetStaticInfo(engine *gin.Engine, appConfigInfo config.AppConfig) {
	engine.Static("/static", appConfigInfo.App.Server.StaticPath)
}
