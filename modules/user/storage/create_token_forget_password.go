package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) CreateTokenForgotPassword(ctx context.Context, data *model.ForgetPassword) error {
	return nil
}
