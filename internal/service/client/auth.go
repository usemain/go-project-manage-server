package client

import (
	"gin-choes-server/api/v1/client/auth"
)

type IAuth interface {
	Code(v auth.CodeRequest) (err error)
	Login(v auth.LoginRequest) (data *auth.LoginResponse, err error)
	Register(v auth.RegisterRequest) (err error)
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
