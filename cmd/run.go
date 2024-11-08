package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"go-project-manage-server/internal/constants"
	"go-project-manage-server/internal/handler"
	"go-project-manage-server/pkg/response"
	"net/http"
	"time"
)

var limiter *ratelimit.Bucket

func init() {
	limiter = ratelimit.NewBucketWithQuantum(10*time.Second, 20, 10)
}

// RateLimiter 请求频繁拦截
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if limiter.TakeAvailable(1) <= 0 {
			response.Result(c, constants.StatusLimiterError, "请勿频繁请求")
			c.Abort()
			return
		}
		c.Next()
	}
}

func Run(port string) {
	app := gin.Default()

	app.Use(RateLimiter())

	app.StaticFS("static", http.Dir("static"))

	api := app.Group("api")
	handler.Auth(api)
	handler.User(api)
	handler.Task(api)

	if err := app.Run(port); err != nil {
		panic("Server run error -> " + err.Error())
	}
}
