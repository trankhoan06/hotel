package biz

import (
	"context"
	"errors"
	"time"
)

func (biz *UserBiz) NewVerifyCode(ctx context.Context, code int, token string, userId int) error {
	now := time.Now()
	now = now.Add(-7 * time.Hour)
	sender, err := biz.store.GetSendCode(ctx, map[string]interface{}{"token": token, "user_id": userId})
	if err != nil {
		return err
	}
	if now.After(sender.Expire) {
		return errors.New("code expired")
	}
	if code != sender.Code {
		return errors.New("code error")
	}
	return nil
}
