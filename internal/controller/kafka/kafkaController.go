package kafkaController

import (
	kafkaService "gKafkaAdmin/internal/apiserver/service/kafka"
	responseRlt "gKafkaAdmin/internal/module"

	"github.com/gin-gonic/gin"
)

func GetTopicList(ctx *gin.Context) {
	responseRlt.Success(ctx, kafkaService.ListAllTopic("123"))
}
