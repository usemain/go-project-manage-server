package service

import (
	"go-project-manage-server/api/user"
)

type IUser interface {
	Userinfo(Uid string) (data *userApi.UserinfoResponse, err error)
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
