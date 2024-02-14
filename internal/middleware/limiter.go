package middleware

import (
	"gin-choes-server/internal/consts"
	"gin-choes-server/utility"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
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
			utility.H(c, consts.StatusLimiterError, "请勿频繁请求")
			c.Abort()
			return
		}
		c.Next()
	}
}
