package biz

import (
	"context"
	TokenProvider "main.go/component"
	"main.go/modules/user/model"
)

type UserStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	Register(ctx context.Context, data *model.CreateUser) error
	DeletedUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, data *model.UpdateUser) error
	VerifyEmail(ctx context.Context, id int) error
	CreateCode(ctx context.Context, data *model.CreateSendCode) error
	GetSendCode(ctx context.Context, cond map[string]interface{}) (*model.SendCode, error)
	ChangePassword(ctx context.Context, password, email string) error
	CreateTokenForgotPassword(ctx context.Context, data *model.ForgetPassword) error
	UpdateTokenForgot(ctx context.Context, token string, user int) error
}
type UserBiz struct {
	store UserStorage
}
type RegisterUser struct {
	store UserStorage
	hash  Hasher
}

func NewRegisterUser(store UserStorage, hasher Hasher) *RegisterUser {
	return &RegisterUser{
		store: store,
		hash:  hasher,
	}
}

type LoginUser struct {
	store    UserStorage
	hash     Hasher
	provider TokenProvider.Provider
}

func NewLoginUser(store UserStorage, provide TokenProvider.Provider, hasher Hasher) *LoginUser {
	return &LoginUser{
		store:    store,
		provider: provide,
		hash:     hasher,
	}
}

type Hasher interface {
	Hash(str string) string
}

func NewUserBiz(store UserStorage) *UserBiz {
	return &UserBiz{store}
}
