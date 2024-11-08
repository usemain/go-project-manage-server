package logic

import (
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"go-project-manage-server/api/auth"
	"go-project-manage-server/internal/constants"
	"go-project-manage-server/internal/global"
	"go-project-manage-server/internal/model"
	"go-project-manage-server/internal/service"
	"go-project-manage-server/pkg/utils"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type IAuth struct{}

func init() {
	service.RegisterServiceAuth(&IAuth{})
}

func (s *IAuth) Code(v authApi.CodeRequest) error {
	sendEmail := v.Email + "_code"
	do, err := global.GVA_REDIS.Do(global.GVA_CTX, "get", sendEmail).Result()
	if err != nil || do != nil {
		return errors.New("发送失败")
	}

	// 生成6位数验证码
	randCode := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	vCode := fmt.Sprintf("%06d", randCode)

	// 发送邮件
	e := email.NewEmail()
	e.From = constants.SEND_EMAIL_USER
	e.To = []string{v.Email}
	e.Subject = "验证码"
	e.HTML = []byte(fmt.Sprintf(constants.SEND_EMAIL_CODE_TEMPLATE, time.Now().Format("2006-01-02 15:04:05"), vCode))
	ePlainAuth := smtp.PlainAuth("", constants.SEND_EMAIL_USER, constants.SEND_EMAIL_PWD, constants.SEND_EMAIL_HOST)
	err = e.Send("smtp.163.com:25", ePlainAuth)
	if err != nil {
		return errors.New("验证码发送失败")
	}

	// 将验证码存入 Redis，设置过期时间5分钟
	_, err = global.GVA_REDIS.SetNX(global.GVA_CTX, sendEmail, vCode, time.Minute*5).Result()
	if err != nil {
		return errors.New("系统错误")
	}

	return nil
}

func (s *IAuth) Login(v authApi.LoginRequest) (data *authApi.LoginResponse, err error) {
	u := &model.USER{}
	if tx := global.GVA_DB.Where("email = ? and status = ?", v.Email, true).Take(u); tx.RowsAffected == 0 {
		return nil, errors.New("用户不存在或已被禁用")
	}

	if valid := utils.ValidPassword(v.Pwd, constants.SECRET, u.Pwd); !valid {
		return nil, errors.New("密码错误")
	}

	token, err := utils.MakeToken(u.Uid, v.Email)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	global.GVA_REDIS.Do(global.GVA_CTX, "set", u.Email+"_token", token)

	var lastTime = time.Now().Format("2006-01-02 15:04:05")
	global.GVA_DB.Model(&model.USER{}).Where("uid = ?", u.Uid).Update("last_time", lastTime)

	// 判断是否组长或者有没有发布任务的权限的组员
	g := &model.GROUP{}
	var isGroupCreateTask = false
	if tx := global.GVA_DB.Where("gid = ?", u.Gid).Take(g); tx.RowsAffected != 0 && u.Gid != 0 {
		// 判断用户是否在组内
		for _, v := range strings.Split(g.GroupUsers, "-") {
			if v == u.Uid {
				isGroupCreateTask = true
				break
			}
		}
	}

	return &authApi.LoginResponse{
		Gid:               u.Gid,
		IsGroupLeader:     g.Uid == u.Uid,
		IsGroupCreateTask: isGroupCreateTask,
		LastTime:          lastTime,
		Email:             u.Email,
		Head:              u.Head,
		Name:              u.Name,
		Sex:               u.Sex,
		VipLevel:          u.VipLevel,
		VipExpireTime:     u.VipExpireTime,
		Secret:            u.Secret,
		Token:             token,
	}, nil
}

func (s *IAuth) Register(v authApi.RegisterRequest) (err error) {
	if tx := global.GVA_DB.Where("email = ?", v.Email).Take(&model.USER{}); tx.RowsAffected != 0 {
		return errors.New("账号已注册")
	}

	if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", v.Email+"_code"); do.Val() != v.Code {
		return errors.New("验证码错误")
	}

	if tx := global.GVA_DB.Create(model.USER{
		Uid:        utils.GenerateUniqueID(12, v.Email),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Email:      v.Email,
		Pwd:        utils.MakePassword(v.Pwd, constants.SECRET),
		Head:       "init.jpg",
		Name:       "新用户" + strconv.Itoa(rand.Int())[0:8],
		Sex:        1,
		VipLevel:   1,
		Secret:     utils.GenerateUniqueID(16, v.Email),
		Status:     true,
	}); tx.RowsAffected == 0 {
		return errors.New("注册失败")
	}

	global.GVA_REDIS.Do(global.GVA_CTX, "del", v.Email+"_code")
	return nil
}
