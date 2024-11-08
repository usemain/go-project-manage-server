package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"go-project-manage-server/internal/constants"
	"time"
)

type MyClaims struct {
	Uid   string `json:"uid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func MakeToken(Uid, email string) (tokenString string, err error) {
	claim := MyClaims{
		Uid:   Uid,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(720 * time.Hour * time.Duration(1))), // 过期时间30天
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                         // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                         // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	tokenString, err = token.SignedString([]byte(constants.SECRET))
	return tokenString, err
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.SECRET), nil // secret
	}
}

func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, Secret())
	if err != nil {
		var ve *jwt.ValidationError
		if errors.As(err, &ve) {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("couldn't handle this token")
}
