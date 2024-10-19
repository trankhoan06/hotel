package biz

import (
	"context"
	"errors"
	"main.go/modules/user/model"
	"time"
)

func (biz *UserBiz) NewVerifyCodeForgot(ctx context.Context, code int, token string, userId int, expire int) (*model.IsEmailToken, error) {
	sendCode, err := biz.store.GetSendCode(ctx, map[string]interface{}{"token": token, "user_id": userId})
	if err != nil {
		return nil, err
	}
	now := time.Now().Add(-7 * time.Hour)
	if now.After(sendCode.Expire) {
		return nil, errors.New("code is expire")
	}
	if sendCode.Code != code {
		return nil, errors.New("code is invalid")
	}
	now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.VerifyCodeForgot(ctx, map[string]interface{}{"token": token, "user_id": userId}, map[string]interface{}{"verify": true, "expire": now}); err != nil {
		return nil, err
	}
	var IsEmail model.IsEmailToken
	IsEmail.Token = token
	IsEmail.UserId = userId
	return &IsEmail, nil
}
