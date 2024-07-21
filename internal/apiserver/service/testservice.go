package service

import (
	responseRlt "gKafkaAdmin/internal/module"
	"gKafkaAdmin/internal/module/vo"
	"gKafkaAdmin/internal/zlog"

	"github.com/gin-gonic/gin"
)

func TestFunc(ctx *gin.Context) {
	var testData vo.TestData
	testData.Description = "123"
	zlog.Info(testData.Description)
	responseRlt.Success(ctx, testData)
}
