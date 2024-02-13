package client

import (
	"gin-choes-server/api/v1/client/auth"
	"gin-choes-server/internal/consts"
	"gin-choes-server/internal/service/client"
	"gin-choes-server/utility"
	"github.com/gin-gonic/gin"
	"regexp"
)

type AuthControllerV1 struct{}

func AuthNewV1() auth.IAuth {
	return &AuthControllerV1{}
}

// Code 获取验证码
func (u *AuthControllerV1) Code(c *gin.Context) {
	v := auth.CodeRequest{}
	if err := c.ShouldBindQuery(&v); err != nil {
		utility.H(c, consts.StatusError, "获取失败")
		return
	}

	if err := client.Auth().Code(v); err == nil {
		utility.H(c, consts.StatusOK, "获取成功")
	} else {
		utility.H(c, consts.StatusError, err.Error())
	}
}

// Login 用户登录
func (u *AuthControllerV1) Login(c *gin.Context) {
	v := auth.LoginRequest{}
	if err := c.ShouldBindJSON(&v); err != nil {
		utility.H(c, consts.StatusError, "登录失败")
		return
	}

	if p := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{5,15}$`); !p.MatchString(v.Pwd) {
		utility.H(c, consts.StatusError, "登录失败")
		return
	}

	if data, err := client.Auth().Login(v); err == nil {
		utility.D(c, consts.StatusOK, "登录成功", data)
	} else {
		utility.H(c, consts.StatusError, "登录失败")
	}
}

// Register 用户注册
func (u *AuthControllerV1) Register(c *gin.Context) {
	v := auth.RegisterRequest{}
	if err := c.ShouldBindJSON(&v); err != nil {
		utility.H(c, consts.StatusError, "注册失败")
		return
	}

	if p := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{5,15}$`); !p.MatchString(v.Pwd) {
		utility.H(c, consts.StatusError, "注册失败")
		return
	}

	if err := client.Auth().Register(v); err == nil {
		utility.H(c, consts.StatusOK, "注册成功")
	} else {
		utility.H(c, consts.StatusError, "注册失败")
	}
}
