package logic

import (
	"errors"
	"go-project-manage-server/api/user"
	"go-project-manage-server/internal/global"
	"go-project-manage-server/internal/model"
	"go-project-manage-server/internal/service"
)

type IUser struct{}

func init() {
	service.RegisterServiceUser(&IUser{})
}

func (s *IUser) Userinfo(uid string) (data *userApi.UserinfoResponse, err error) {
	u := &model.USER{}
	if err := global.GVA_DB.Where("uid = ?", uid).Take(u); err.RowsAffected == 0 {
		return nil, errors.New("获取失败")
	}

	return &userApi.UserinfoResponse{
		Email:         u.Email,
		Head:          u.Head,
		Name:          u.Name,
		Sex:           u.Sex,
		VipLevel:      u.VipLevel,
		VipExpireTime: u.VipExpireTime,
		Secret:        u.Secret,
		Status:        u.Status,
	}, nil
}
