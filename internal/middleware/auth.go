package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/internal/constants"
	"go-project-manage-server/internal/global"
	"go-project-manage-server/pkg/response"
	"go-project-manage-server/pkg/utils"
	"strings"
)

// AuthToken 校验Token
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			response.ResultAuthError(c, constants.StatusUnauthorized, "身份认证失败")
			c.Abort()
			return
		}

		token := c.GetHeader("Authorization")[strings.LastIndex(c.GetHeader("Authorization"), "Bearer ")+7:]
		data, err := utils.ParseToken(token)
		if err != nil {
			response.ResultAuthError(c, constants.StatusUnauthorized, "身份认证失败")
			c.Abort()
			return
		}

		if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", data.Email+"_token"); do.Val() != token {
			response.ResultAuthError(c, constants.StatusUnauthorized, "身份认证失败")
			c.Abort()
			return
		} else {
			c.Set("uid", data.Uid)
			c.Next()
		}
	}
}
