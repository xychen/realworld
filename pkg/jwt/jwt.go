package jwt

import (
	"github.com/golang-jwt/jwt"
	"realworld/utils"
	"time"
)

const AppKey = "key"
const AppSecret = "cxy"

const GlobalSecret = "realworld"
const GlobalExpires = 7200
const GlobalIssuer = "cxy"

func GetJWTSecret() []byte {
	return []byte(GlobalSecret)
}

// GenerateToken 生成签名.
func GenerateToken(appKey, appSecret, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(GlobalExpires * time.Second)
	claims := Claims{
		AppKey:    utils.MD5([]byte(AppKey)),
		AppSecret: utils.MD5([]byte(AppSecret)),
		UserName:  userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    GlobalIssuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

// Claims (请求权).
type Claims struct {
	AppKey    string
	AppSecret string
	UserName  string
	jwt.StandardClaims
}

// ParseToken 解析签名.
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
