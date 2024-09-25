package biz

import (
	"context"
	"errors"
	"fmt"
	"main.go/common"
	"main.go/modules/user/model"
	"time"
)

func (biz *UserBiz) NewVerifyCode(ctx context.Context, code int, token string, userId int, expire int) (string, error) {
	now := time.Now()
	now = now.Add(-7 * time.Hour)
	sender, err := biz.store.GetSendCode(ctx, map[string]interface{}{"token": token, "user_id": userId})
	if err != nil {
		return "", err
	}
	if now.After(sender.Expire) {
		return "", errors.New("code expired")
	}
	if code != sender.Code {
		return "", errors.New("code error")
	}
	token1 := common.GetSalt(30)
	var data model.ForgetPassword
	data.UserId = userId
	data.Token = token1
	now = now.Add(time.Duration(expire) * time.Second)
	data.Expire = now
	go func() {
		errToken := biz.store.CreateTokenForgotPassword(ctx, &data)
		fmt.Println(errToken)
	}()
	return token1, nil
}
