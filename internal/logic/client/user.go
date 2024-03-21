package client

import (
	"errors"
	"gin-choes-server/api/v1/client/user"
	"gin-choes-server/internal/global"
	"gin-choes-server/internal/model"
	"gin-choes-server/internal/service/client"
)

type SUser struct{}

func init() {
	client.RegisterServiceUser(&SUser{})
}

func (s *SUser) Userinfo(uid string) (data *user.UserinfoResponse, err error) {
	u := &model.USER{}
	if err := global.GVA_DB.Where("uid = ?", uid).Take(u); err.RowsAffected == 0 {
		return nil, errors.New("获取失败")
	}

	return &user.UserinfoResponse{
		Email:         u.Email,
		Head:          u.Head,
		Name:          u.Name,
		Sex:           u.Sex,
		VipLevel:      u.VipLevel,
		VipExpireTime: u.VipExpireTime,
		Status:        u.Status,
	}, nil
}
