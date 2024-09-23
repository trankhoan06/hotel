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
	var eto model.IsEmailToken
	token := common.GetSalt(30)
	go func() {
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
		_ = biz.store.CreateCode(ctx, sendCode)
		_ = email.SendCode(email1, code)
	}()
	eto.IsEmail = user.IsEmail
	eto.Token = token
	eto.UserId = user.Id
	return &eto, nil
}
