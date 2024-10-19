package biz

import "context"

func (biz *RegisterUser) NewChangePassword(ctx context.Context, password string, userId int) error {
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
