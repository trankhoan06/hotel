package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/email"
	"main.go/modules/user/model"
	"time"
)

func (r *RegisterUser) NewRegister(ctx context.Context, data *model.CreateUser) (string, error) {
	if _, err := r.store.FindUser(ctx, map[string]interface{}{"email": data.Email}); err == nil {
		return "", errors.New("email has been register")
	}
	data.Salt = common.GetSalt(50)
	data.Password = r.hash.Hash(data.Salt + data.Password)
	if errUser := r.store.Register(ctx, data); errUser != nil {
		return "", errUser
	}
	token := common.GetSalt(30)
	go func() {
		code := common.GenerateRandomCode()
		now := time.Now()
		now = now.Add(-7 * time.Hour)
		now = now.Add(1 * time.Minute)
		sendCode := &model.CreateSendCode{
			Code:   code,
			UserId: *data.Id,
			Expire: now,
			Token:  token,
		}
		_ = r.store.CreateCode(ctx, sendCode)
		_ = email.SendCode(data.Email, code)
	}()

	return token, nil
}
