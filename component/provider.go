package TokenProvider

import "main.go/modules/user/model"

type Provider interface {
	General(data Payload, expire int) (Token, error)
	Validate(token string) (Payload, error)
	GetSecret() string
}
type Payload interface {
	GetUser() int
	GetRole() *model.RoleUser
}
type Token interface {
	GetToken() string
}
