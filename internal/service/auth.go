package service

import (
	"go-project-manage-server/api/auth"
)

type IAuth interface {
	Code(v authApi.CodeRequest) (err error)
	Login(v authApi.LoginRequest) (data *authApi.LoginResponse, err error)
	Register(v authApi.RegisterRequest) (err error)
}

var serviceAuth IAuth

func RegisterServiceAuth(s IAuth) {
	serviceAuth = s
}

func Auth() IAuth {
	if serviceUser == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return serviceAuth
}
