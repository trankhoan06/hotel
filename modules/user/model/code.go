package model

import "time"

type SendCode struct {
	UserId   int       `json:"user_id" gorm:"column:user_id"`
	Code     int       `json:"code" gorm:"column:code"`
	Token    string    `json:"token" gorm:"column:token"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	Expire   time.Time `json:"expire" gorm:"column:expire"`
}

type CreateSendCode struct {
	UserId int       `json:"user_id" gorm:"column:user_id"`
	Code   int       `json:"code" gorm:"column:code"`
	Token  string    `json:"token" gorm:"column:token"`
	Expire time.Time `json:"expire" gorm:"column:expire"`
}

func (SendCode) TableName() string       { return "send_code" }
func (CreateSendCode) TableName() string { return "send_code" }
