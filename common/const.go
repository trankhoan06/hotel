package common

type Payload struct {
	URole string `json:"role"`
	UId   int    `json:"user_id"`
}

func (p *Payload) GetRole() string { return p.URole }
func (p *Payload) GetUser() int    { return p.UId }

type Requester interface {
	GetUserId() int
	GetRole() string
	GetEmail() string
}
