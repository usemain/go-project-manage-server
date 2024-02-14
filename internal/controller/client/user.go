package client

import (
	"gin-choes-server/api/v1/client/user"
	"gin-choes-server/internal/consts"
	"gin-choes-server/internal/service/client"
	"gin-choes-server/utility"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserControllerV1 struct{}

func UserNewV1() user.IUser {
	return &UserControllerV1{}
}

// Userinfo 用户信息
func (u *UserControllerV1) Userinfo(c *gin.Context) {
	uid, ok := c.Get("uid")
	if !ok {
		utility.H(c, consts.StatusError, "获取失败")
	}

	if data, err := client.User().Userinfo(uid.(string)); err == nil {
		token := c.GetHeader("Authorization")[strings.LastIndex(c.GetHeader("Authorization"), "Bearer ")+7:]
		data.Token = token
		utility.D(c, consts.StatusOK, "获取成功", data)
	} else {
		utility.H(c, consts.StatusError, "获取失败")
	}
}
