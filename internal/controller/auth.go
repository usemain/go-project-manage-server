package controller

import (
	"github.com/gin-gonic/gin"
	"go-project-manage-server/api/auth"
	"go-project-manage-server/internal/constants"
	"go-project-manage-server/internal/service"
	"go-project-manage-server/pkg/response"
	"regexp"
)

type AuthController struct{}

func AuthNew() authApi.IAuth {
	return &AuthController{}
}

// Code 获取验证码
func (u *AuthController) Code(c *gin.Context) {
	v := authApi.CodeRequest{}
	if err := c.ShouldBindQuery(&v); err != nil {
		response.Result(c, constants.StatusError, "请求参数不合法")
		return
	}

	if err := service.Auth().Code(v); err != nil {
		response.Result(c, constants.StatusError, err.Error())
	} else {
		response.Result(c, constants.StatusOK, "获取成功")
	}
}

// Login 用户登录
func (u *AuthController) Login(c *gin.Context) {
	v := authApi.LoginRequest{}
	if err := c.ShouldBindJSON(&v); err != nil {
		response.Result(c, constants.StatusError, "请求参数不合法")
		return
	}

	if p := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{5,15}$`); !p.MatchString(v.Pwd) {
		response.Result(c, constants.StatusError, "账号或密码格式不符合要求")
		return
	}

	if data, err := service.Auth().Login(v); err != nil {
		response.Result(c, constants.StatusError, err.Error())
	} else {
		response.ResultData(c, constants.StatusOK, "登录成功", data)
	}
}

// Register 用户注册
func (u *AuthController) Register(c *gin.Context) {
	v := authApi.RegisterRequest{}
	if err := c.ShouldBindJSON(&v); err != nil {
		response.Result(c, constants.StatusError, "请求参数不合法")
		return
	}

	if p := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]{5,15}$`); !p.MatchString(v.Pwd) {
		response.Result(c, constants.StatusError, "账号或密码格式不符合要求")
		return
	}

	if err := service.Auth().Register(v); err != nil {
		response.Result(c, constants.StatusError, err.Error())
	} else {
		response.Result(c, constants.StatusOK, "注册成功")
	}
}
