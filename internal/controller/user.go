package controller

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/api/user"
	"go-project-manage-server/internal/constants"
	"go-project-manage-server/internal/service"
	"go-project-manage-server/pkg/response"
	"strings"
)

type UserController struct{}

func UserNew() userApi.IUser {
	return &UserController{}
}

// Userinfo 用户信息
func (u *UserController) Userinfo(c *gin.Context) {
	uid, ok := c.Get("uid")
	if !ok {
		response.Result(c, constants.StatusError, "获取失败")
	}

	if data, err := service.User().Userinfo(uid.(string)); err != nil {
		response.Result(c, constants.StatusError, err.Error())
	} else {
		token := c.GetHeader("Authorization")[strings.LastIndex(c.GetHeader("Authorization"), "Bearer ")+7:]
		data.Token = token
		response.ResultData(c, constants.StatusOK, "获取成功", data)
	}
}
