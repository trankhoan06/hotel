package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"main.go/common"
	TokenProvider "main.go/component"
	"time"
)

type JwtProvider struct {
	Prefix string
	Secret string
}

func NewJwtProvider(Prefix string, Secret string) *JwtProvider {
	return &JwtProvider{Prefix: Prefix, Secret: Secret}
}

type token struct {
	Token  string    `json:"token"`
	Expire int       `json:"expire"`
	Create time.Time `json:"create"`
}

func (t *token) GetToken() string {
	return t.Token
}
func (j *JwtProvider) GetSecret() string {
	return j.Secret
}

type MyClaim struct {
	Payload common.Payload `json:"payload"`
	jwt.StandardClaims
}

func (j *JwtProvider) General(data TokenProvider.Payload, expire int) (TokenProvider.Token, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaim{
		common.Payload{
			URole: data.GetRole(),
			UId:   data.GetUser(),
		},
		jwt.StandardClaims{
			Id:        fmt.Sprint(now.UnixNano()),
			ExpiresAt: now.Add(time.Duration(expire) * time.Second).Unix(),
			IssuedAt:  now.Unix(),
		},
	})
	to, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}
	return &token{
		Token:  to,
		Expire: expire,
		Create: now,
	}, nil
}
func (j *JwtProvider) Validate(token string) (TokenProvider.Payload, error) {
	myToken, err := jwt.ParseWithClaims(token, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if !myToken.Valid {
		return nil, errors.New("token is invalid")
	}
	to, ok := myToken.Claims.(*MyClaim)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	return &to.Payload, nil
}
