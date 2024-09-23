package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) CreateCode(ctx context.Context, data *model.CreateSendCode) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
