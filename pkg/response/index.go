package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Result(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func ResultData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func ResultDataList(c *gin.Context, code int, msg string, data interface{}, count int64) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"data":  data,
		"count": count,
	})
}

func ResultAuthError(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"code": code,
		"msg":  msg,
	})
}
