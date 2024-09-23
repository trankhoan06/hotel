package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	var data model.User
	if err := s.db.Table("user").Where(cond).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
