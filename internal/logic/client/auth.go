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
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")
	// 设置邮件主题
	e.Subject = "邮箱验证码"
	// 设置文件发送的内容
	content := fmt.Sprintf(`
	<div>
        <div>
            尊敬的%s，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>您于 %s 提交的邮箱验证，本次验证码为
				<br/>
				<br/>
				<span style="padding: 10px 30px;background-color: rgb(161, 162, 162);margin: 0 10px;border-radius: 10px;">
					%s
				</span>
				<br/>
				<br/>
				为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。
			</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>
	`, e.To[0], t, vCode)
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
	if err := global.GVA_DB.Where("email = ?", v.Email).Take(u); err.RowsAffected == 0 {
		return nil, errors.New("用户不存在")
	}

	if valid := utility.ValidPassword(v.Pwd, consts.SECRET, u.Pwd); !valid {
		return nil, errors.New("密码错误")
	}

	token, err := utility.MakeToken(u.UID)
	if err != nil {
		return nil, errors.New("生成Token失败")
	}

	global.GVA_REDIS.Do(global.GVA_CTX, "set", u.UID+"_token", token)

	return &auth.LoginResponse{
		Email:  u.Email,
		Head:   u.Head,
		Name:   u.Name,
		Sex:    u.Sex,
		Level:  u.Level,
		Status: u.Status,
		Token:  token,
	}, nil
}

// Register 用户注册
func (s *SAuth) Register(v auth.RegisterRequest) (err error) {
	if err := global.GVA_DB.Where("email = ?", v.Email).Take(&model.USER{}); err.RowsAffected != 0 {
		return errors.New("账号已注册")
	}

	if do := global.GVA_REDIS.Do(global.GVA_CTX, "get", v.Email+"_code"); do.Val() != v.Code {
		return errors.New("验证码错误")
	}

	if err := global.GVA_DB.Create(model.USER{
		UID:    utility.GenerateUniqueID(32),
		Email:  v.Email,
		Pwd:    utility.MakePassword(v.Pwd, consts.SECRET),
		Head:   "init.jpg",
		Name:   "新用户" + strconv.Itoa(rand.Int())[0:8],
		Sex:    1,
		Level:  1,
		Status: 1,
	}); err.RowsAffected == 0 {
		return errors.New("注册失败")
	}

	global.GVA_REDIS.Do(global.GVA_CTX, "del", v.Email+"_code")
	return nil
}
