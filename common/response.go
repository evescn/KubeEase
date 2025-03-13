package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func ResponseParamInvalid(c *gin.Context, err error) {
	zap.L().Error(fmt.Sprintf("绑定参数失败， %v\n", err))
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  fmt.Sprintf("绑定参数失败， %v\n", err),
		"data": nil,
	})
}

func ResponseFailed(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"msg":  err.Error(),
		"data": nil,
	})
}

func ResponseOk(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}
