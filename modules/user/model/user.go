package model

import "time"

type StatusUser int
type RoleUser int

const (
	RoleUserHost RoleUser = iota
	RoleUserUser
)
const (
	StatusUserDeleted StatusUser = iota
	StatusUserDoing
)

type User struct {
	Id        int         `json:"id" gorm:"column:id"`
	Role      *RoleUser   `json:"role" gorm:"column:role"`
	Salt      string      `json:"salt" gorm:"column:salt"`
	FirstName string      `json:"first_name" gorm:"column:first_name"`
	LastName  string      `json:"last_name" gorm:"column:last_name"`
	Phone     string      `json:"phone" gorm:"column:phone"`
	Email     string      `json:"email" gorm:"column:email"`
	Status    *StatusUser `json:"status" gorm:"column:status"`
	Password  string      `json:"password" gorm:"column:password"`
	IsEmail   bool        `json:"is_email" gorm:"column:is_email"`
	CreateAt  time.Time   `json:"createAt" gorm:"column:create_at"`
	UpdateAt  time.Time   `json:"updateAt" gorm:"column: update_at"`
}
type CreateUser struct {
	Id        *int   `json:"id" gorm:"column:id"`
	Salt      string `json:"-" gorm:"column:salt"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
}
type UpdateUser struct {
	Email     string  `json:"-" gorm:"column:email"`
	FirstName *string `json:"first_name" gorm:"column:first_name"`
	LastName  *string `json:"last_name" gorm:"column:last_name"`
	Phone     *string `json:"phone" gorm:"column:phone"`
}
type Login struct {
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
type IsEmailToken struct {
	IsEmail bool   `json:"is_email" gorm:"column:is_email"`
	Token   string `json:"token" gorm:"column:token"`
	UserId  int    `json:"user_id" gorm:"column:user_id"`
}

func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetEmail() string {
	return u.Email
}
func (u *User) GetRole() *RoleUser {
	return u.Role
}
func (CreateUser) TableName() string {
	return "user"
}
func (User) TableName() string {
	return "user"
}
func (UpdateUser) TableName() string {
	return "user"
}
func (Login) TableName() string {
	return "user"
}
