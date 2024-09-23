package biz

import (
	"context"
	"errors"
	"main.go/common"
	"main.go/modules/user/model"
	"time"
)

func (biz *LoginUser) VerifyCodeEmail(ctx context.Context, token string, userId int, expire int, code int) (string, error) {
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
	if err := biz.store.VerifyEmail(ctx, userId); err != nil {
		return "", err
	}
	roleUserPtr := model.RoleUserUser
	var payLoad = &common.Payload{
		UId:   userId,
		URole: &roleUserPtr,
	}
	token1, err := biz.provider.General(payLoad, expire)
	if err != nil {
		return "", err
	}
	return token1.GetToken(), nil
}
