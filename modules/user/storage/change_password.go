package storage

import "context"

func (s *SqlModel) ChangePassword(ctx context.Context, password, email string) error {
	if err := s.db.Table("user").Where("email=?", email).Update("password", password).Error; err != nil {
		return err
	}
	return nil
}
