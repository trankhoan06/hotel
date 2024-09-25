package storage

import (
	"context"
	"time"
)

func (s *SqlModel) UpdateTokenForgot(ctx context.Context, token string, user int) error {
	now := time.Now()
	now = now.Add(-7 * time.Hour)
	if err := s.db.Table("token_forgot_pasword").Where("token=? && user_id=?", token, user).Update("expire", now).Error; err != nil {
		return err
	}
	return nil
}
