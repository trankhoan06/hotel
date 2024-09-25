package biz

import (
	"context"
	"fmt"
)

func (biz *RegisterUser) ChangePasswordForgot(ctx context.Context, token, password, email string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return err
	}
	password = biz.hash.Hash(user.Salt + password)
	if err := biz.store.ChangePassword(ctx, password, email); err != nil {
		return err
	}
	go func() {
		if err := biz.store.UpdateTokenForgot(ctx, token, user.Id); err != nil {
			fmt.Println(err)
		}
	}()
	return nil
}
