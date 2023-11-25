package common

type Requester interface {
	GetID() uint
	GetUsername() string
	GetEmail() string
	GetRole() string
}
