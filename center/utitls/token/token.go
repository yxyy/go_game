package token

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"lhc.go.game.center/model"
	"strconv"
	"time"
)
//标准中的保留字如下：
//变量名	英文全写	备注
//iss	Issuer	该JWT的发布者
//sub	Subject	该JWT面向的主体或者用户
//aud	Audience	接收该JWT的用户
//exp	Expiration Time	过期时间（单位为秒）
//nbf	Not Before	开始时间（单位为秒），在该时间之前无效
//iat	Issued At	发布时间（单位为秒）
//jti	JWT ID	JWT唯一标识，区分不同发布者的统一的标识

// 生成 access_token
func CreateAccessToken(user *model.User) (string,error) {
	exp , err := strconv.Atoi(viper.GetString("access_token_exp"))
	if err != nil {
		exp = 2
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id" : user.Id,
		"account" : user.Account,
		"nickname" : user.Nickname,
		"exp" : time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
	})
	token, err := claims.SignedString([]byte(viper.GetString("secret_key")))
	if err!=nil{
		return "" , err
	}
	return token , nil
}

//生成 refresh_token
func CreateRefreshToken(user *model.User) (string,error) {
	exp , err := strconv.Atoi(viper.GetString("refresh_token_exp"))
	if err != nil {
		exp = 30
	}
	exp_time := time.Now().Add(time.Hour * 24 * time.Duration(exp)).Unix()
	sign :=fmt.Sprint(md5.Sum([]byte(user.Nickname + user.Account + strconv.Itoa(user.Id) + strconv.Itoa(int(exp_time)))))
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"id" : user.Id,
		"account" : user.Account,
		"nickname" : user.Nickname,
		"exp" : time.Now().Add(time.Hour * 24 * 30).Unix(),
		"sign" : sign,
	})
	token, err := claims.SignedString([]byte(viper.GetString("secret_key")))
	if err!=nil{
		return "" , err
	}
	return token , nil
}

func secret() jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("secret_key")),nil
	}
}

// 全部解析验证，完全解析验证才有数据
func ParseToken(token string) (*jwt.Token , error) {
	return jwt.Parse(token, secret())
}

type JWTClaims struct {  // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	UserID      uint     `json:"id"`
	NickName    string   `json:"nickname"`
	Account    string   `json:"account"`
	Permissions []string `json:"permissions"`
}

// 只解析 不验证
func ParseAndSignatureToken (token string) (*jwt.Token , error)  {
	return jwt.ParseWithClaims(token,&JWTClaims{},secret())
}


