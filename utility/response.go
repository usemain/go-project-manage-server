package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func H(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

func D(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

func L(c *gin.Context, code int, msg string, data interface{}, total int) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"data":  data,
		"total": total,
	})
}
