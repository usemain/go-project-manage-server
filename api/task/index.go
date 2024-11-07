package taskApi

import (
	"github.com/gin-gonic/gin"
)

type ITask interface {
	SelectDateTask(c *gin.Context)
}
