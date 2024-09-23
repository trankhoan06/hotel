package storage

import (
	"context"
	"main.go/modules/user/model"
)

func (s *SqlModel) GetSendCode(ctx context.Context, cond map[string]interface{}) (*model.SendCode, error) {
	var data model.SendCode
	if err := s.db.Where(cond).Last(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
