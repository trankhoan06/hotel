package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (biz *LoginUser) NewLogin(ctx context.Context, login *model.Login, expire int) (*model.IsEmailToken, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": login.Email})
	if err != nil {
		return nil, err
	}
	login.Password = biz.hash.Hash(user.Salt + login.Password)
	if login.Password != user.Password {
		return nil, errors.New("email of password wrong")
	}
	var Eto model.IsEmailToken
	Eto.IsEmail = user.IsEmail
	Eto.UserId = user.Id
	if !user.IsEmail {
		token := common.GetSalt(30)
		Eto.Token = token
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
		_ = email.SendCode(login.Email, code)
		return &Eto, nil
	}
	payload := &common.Payload{
		UId:   user.Id,
		URole: user.Role,
	}
	token, errToken := biz.provider.General(payload, expire)
	if errToken != nil {
		return nil, err
	}
	Eto.Token = token.GetToken()
	return &Eto, nil
}
