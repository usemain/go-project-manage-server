package client

import (
	"gin-choes-server/api/v1/client/user"
)

type IUser interface {
	Userinfo(Uid string) (data *user.UserinfoResponse, err error)
}

var serviceUser IUser

func RegisterServiceUser(s IUser) {
	serviceUser = s
}

func User() IUser {
	if serviceUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return serviceUser
}
