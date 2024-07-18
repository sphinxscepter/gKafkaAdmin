package service

import (
	"gKafkaAdmin/internal/global"
	"gKafkaAdmin/internal/module/vo"
	"gKafkaAdmin/internal/zlog"

	"github.com/gin-gonic/gin"
)

func TestFunc(ctx *gin.Context) string {
	zlog.Info(ctx.ClientIP())
	zlog.Info(ctx.FullPath())
	zlog.Info(ctx.ContentType())
	zlog.Info(ctx.RemoteIP())
	return "TEST OK"
}

func TestFunc2(ctx *gin.Context) vo.ResultStructure {
	zlog.Info(ctx.ClientIP())
	zlog.Info(ctx.FullPath())
	zlog.Info(ctx.ContentType())
	zlog.Info(ctx.RemoteIP())
	var result vo.ResultStructure
	result.Code = global.HTTP_RESPONSE_SUCESS
	result.Message = "OK2"
	var testData vo.TestData
	testData.Desc = "123"
	result.Data = testData
	return result
}
