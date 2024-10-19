package biz

import (
	"context"
	"main.go/common"
	"main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (biz *LoginUser) ForgotPassword(ctx context.Context, email1 string) (*model.IsEmailToken, error) {
	user, errUser := biz.store.FindUser(ctx, map[string]interface{}{"email": email1})
	if errUser != nil {
		return nil, errUser
	}
	token := common.GetSalt(30)
	code := common.GenerateRandomCode()
	now := time.Now()
	now = now.Add(-7 * time.Hour)
	now = now.Add(1 * time.Minute)
	sendCode := &model.CreateSendCode{
		Code:   code,
		UserId: user.Id,
		Expire: now,
		Token:  token,
	}
	if err := biz.store.CreateCode(ctx, sendCode); err != nil {
		return nil, err
	}
	chanel := make(chan error, 1)
	go func() {
		defer close(chanel)
		if err := email.SendCodeForgot(email1, code); err != nil {
			chanel <- err
		}
	}()
	var eto model.IsEmailToken
	eto.IsEmail = user.IsEmail
	eto.Token = token
	eto.UserId = user.Id
	return &eto, <-chanel
}
