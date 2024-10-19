package biz

import (
	"context"
	"errors"
	"time"
)

func (biz *RegisterUser) ChangePasswordForgot(ctx context.Context, token, password string, userId int) error {
	sendCode, err := biz.store.GetSendCode(ctx, map[string]interface{}{"token": token, "user_id": userId})
	if err != nil {
		return err
	}
	if !sendCode.Verify {
		return errors.New("please verify code")
	}
	now := time.Now().Add(-7 * time.Hour)
	if now.After(sendCode.Expire) {
		return errors.New("token expired")
	}
	if err := biz.store.VerifyCodeForgot(ctx, map[string]interface{}{"user_id": userId, "token": token}, map[string]interface{}{"expire": now}); err != nil {
		return err
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": userId})
	if err != nil {
		return err
	}
	password = biz.hash.Hash(user.Salt + password)
	if err := biz.store.ChangePassword(ctx, password, userId); err != nil {
		return err
	}
	return nil
}
