package storage

import "context"

func (s *SqlModel) ChangePassword(ctx context.Context, password string, userId int) error {
	if err := s.db.Table("user").Where("user_id=?", userId).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
