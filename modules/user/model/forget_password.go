package model

import "time"

type ForgetPassword struct {
	UserId int       `json:"user_id" gorm:"column:user_id"`
	Token  string    `json:"token" gorm:"column:token"`
	Expire time.Time `json:"expire" gorm:"column:expire"`
}

func (ForgetPassword) TableName() string { return "token_forgot_password" }
