package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) VerifyEmail(ctx context.Context, id int) error {
	if err := s.db.Table(model.User{}.TableName()).Where("id=?", id).Update("is_email", 1).Error; err != nil {
		return err
	}
	return nil
}
