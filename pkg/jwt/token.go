package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type (
	TokenOption struct {
		AccessSecret string
		AccessExpire int64
		Fields       map[string]interface{}
	}

	Token struct {
		AccessToken  string `json:"access_token"`
		AccessExpire int64  `json:"access_expire"`
	}
)

func BuildToken(opt *TokenOption) (*Token, error) {
	now := time.Now().Unix()
	token, err := genToken(opt.AccessSecret, now, opt.AccessExpire, opt.Fields)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token,
		AccessExpire: now + opt.AccessExpire,
	}, nil
}

func genToken(secret string, iat, seconds int64, payload map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + seconds
	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secret))
}
