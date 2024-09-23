package common

import "main.go/modules/user/model"

type Payload struct {
	URole *model.RoleUser `json:"role"`
	UId   int             `json:"user_id"`
}

func (p *Payload) GetRole() *model.RoleUser { return p.URole }
func (p *Payload) GetUser() int             { return p.UId }

type Requester interface {
	GetUserId() int
	GetRole() *model.RoleUser
	GetEmail() string
}

const CurrUser = "current_user"
