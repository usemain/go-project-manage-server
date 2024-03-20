package middleware

import (
	"gin-choes-server/internal/consts"
	"gin-choes-server/internal/global"
	"gin-choes-server/utility"
	"github.com/gin-gonic/gin"
	"strings"
)

// AuthToken 校验Token
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			utility.ResultAuthError(c, consts.StatusUnauthorized, "身份认证失败")
			c.Abort()
			return
		}

		token := c.GetHeader("Authorization")[strings.LastIndex(c.GetHeader("Authorization"), "Bearer ")+7:]
		data, err := utility.ParseToken(token)
		if err != nil {
			utility.ResultAuthError(c, consts.StatusUnauthorized, "身份认证失败")
			c.Abort()
			return
		}

		if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", data.UID+"_token"); do.Val() != token {
			utility.ResultAuthError(c, consts.StatusUnauthorized, "身份认证失败")
			c.Abort()
			return
		} else {
			c.Set("uid", data.UID)
			c.Next()
		}
	}
}
