package service

import (
	"adventurer/global"
	"adventurer/model"
	"adventurer/model/request"
	"adventurer/util"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func Register(req request.UserRegisterRequest) error {
	var count uint
	global.DL_DB.Model(&model.User{}).Where("username=? and deleted_at is null", req.Username).Count(&count)

	if count != 0 {
		return errors.New("register failed, user already exists")
	}

	var account model.User
	if err := util.MapTo(&req, &account); err != nil {
		return errors.New("register failed")
	}
	account.CreatedAt = time.Now()
	account.Password = util.MD5V([]byte(account.Password))

	return global.DL_DB.Create(&account).Error
}

func Login(user model.User) (token string, success bool) {
	var err error
	var count uint
	global.DL_DB.Model(&model.User{}).Where("username=? and password=?", user.Username, util.MD5V([]byte(user.Password))).Count(&count)
	if count != 0 {
		if token, err = generateToken(user.Username); err != nil {
			return "", false
		}

		return token, true
	}

	return "", false
}

// 假定这里的username是唯一标识
func generateToken(username string) (string, error) {
	authConfig := global.DL_CONFIG.Auth
	// 创建一个我们自己的声明
	c := model.CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(authConfig.ExpireMinutes)).Unix(), // 过期时间
			Issuer:    authConfig.Issuer,                                                            // 签发人
			Audience:  authConfig.Audience,                                                          //接收jwt的一方
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 把密钥转称byte数组
	secret := []byte(authConfig.Secret)

	// 使用指定的secret并获得完整的编码后的字符串token
	return token.SignedString(secret)
}
