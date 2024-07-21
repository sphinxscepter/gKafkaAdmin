package responseRlt

import (
	"gKafkaAdmin/internal/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RespResultStructure struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, rlt any) {
	ctx.JSON(http.StatusOK, RespResultStructure{
		global.HTTP_RESPONSE_SUCESS,
		"OK",
		rlt,
	})
}

func Error(ctx *gin.Context, errDesc string) {
	ctx.JSON(http.StatusOK, RespResultStructure{
		global.HTTP_RESPONSE_ERROR,
		errDesc,
		nil,
	})
}
