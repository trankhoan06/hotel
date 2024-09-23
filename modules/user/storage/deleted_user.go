package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) DeletedUser(ctx context.Context, id int) error {
	if err := s.db.Table(model.User{}.TableName()).Where("id=?", id).Update("status", model.StatusUserDeleted).Error; err != nil {
		return err
	}
	return nil
}
