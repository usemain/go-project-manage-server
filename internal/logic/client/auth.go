package client

import (
	"errors"
	"fmt"
	"gin-choes-server/api/v1/client/auth"
	"gin-choes-server/internal/consts"
	"gin-choes-server/internal/global"
	"gin-choes-server/internal/model"
	"gin-choes-server/internal/service/client"
	"gin-choes-server/utility"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type SAuth struct{}

func init() {
	client.RegisterServiceAuth(&SAuth{})
}

// Code 获取验证码
func (s *SAuth) Code(v auth.CodeRequest) (err error) {
	if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", v.Email+"_code"); do.Val() != nil {
		return errors.New("验证码不能频繁发送")
	}
	e := email.NewEmail()
	e.From = fmt.Sprintf("l111il@163.com")
	e.To = []string{v.Email}
	// 生成6位随机验证码
	vCode := fmt.Sprintf("%v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	e.Subject = "邮箱验证码"
	content := fmt.Sprintf(`
	<div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>
				您于 %s 提交的邮箱验证，本次验证码为:<span>%s</span>
				<br/>
				为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
			</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>
	`, time.Now().Format("2006-01-02 15:04:05"), vCode)
	e.HTML = []byte(content)
	// 设置服务器相关的配置
	if err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "l111il@163.com", "TMGUTAJMNWXOLSYY", "smtp.163.com")); err != nil {
		return errors.New("发送失败")
	}

	global.GVA_REDIS.Do(global.GVA_CTX, "set", v.Email+"_code", vCode)
	global.GVA_REDIS.Do(global.GVA_CTX, "expire", v.Email+"_code", 300)
	return nil
}

// Login 用户登录
func (s *SAuth) Login(v auth.LoginRequest) (data *auth.LoginResponse, err error) {
	u := &model.USER{}
	if tx := global.GVA_DB.Where("email = ? and status = ?", v.Email, true).Take(u); tx.RowsAffected == 0 {
		return nil, errors.New("用户不存在或已被禁用")
	}

	if valid := utility.ValidPassword(v.Pwd, consts.SECRET, u.Pwd); !valid {
		return nil, errors.New("密码错误")
	}

	token, err := utility.MakeToken(u.Uid, v.Email)
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

	return &auth.LoginResponse{
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

// Register 用户注册
func (s *SAuth) Register(v auth.RegisterRequest) (err error) {
	if tx := global.GVA_DB.Where("email = ?", v.Email).Take(&model.USER{}); tx.RowsAffected != 0 {
		return errors.New("账号已注册")
	}

	if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", v.Email+"_code"); do.Val() != v.Code {
		return errors.New("验证码错误")
	}

	if tx := global.GVA_DB.Create(model.USER{
		Uid:        utility.GenerateUniqueID(18, v.Email),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Email:      v.Email,
		Pwd:        utility.MakePassword(v.Pwd, consts.SECRET),
		Head:       "init.jpg",
		Name:       "新用户" + strconv.Itoa(rand.Int())[0:8],
		Sex:        1,
		VipLevel:   1,
		Secret:     utility.GenerateUniqueID(16, v.Email),
		Status:     true,
	}); tx.RowsAffected == 0 {
		return errors.New("注册失败")
	}

	global.GVA_REDIS.Do(global.GVA_CTX, "del", v.Email+"_code")
	return nil
}
