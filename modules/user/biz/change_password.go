package biz

import "context"

func (biz *RegisterUser) NewChangePassword(ctx context.Context, password string, email string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return err
	}
	password = biz.hash.Hash(user.Salt + password)
	if err := biz.store.ChangePassword(ctx, password, email); err != nil {
		return err
	}
	return nil
}
