package common

import "errors"

var (
	RecordNotFound      = errors.New("record not found")
	RecordIsBlocked     = errors.New("record is blocked")
	InsufficientBalance = errors.New("insufficient balance")
	EntityExists        = errors.New("entity already exists")
	NoPermission        = errors.New("user has no permission")
)
