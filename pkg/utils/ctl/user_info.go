package ctl

import (
	"context"
	"errors"
)

type key int

var userKey key

type UserInfo struct {
	Id uint `json:"id"`
}

// GetUserInfo retrieves user information from the context
func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return nil, errors.New("error retrieving user information")
	}
	return user, nil
}

// NewContext creates a new context with user information
func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// FromContext retrieves user information from the context
func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

// InitUserInfo initializes user information in the context (currently empty)
func InitUserInfo(ctx context.Context) {
	// TODO: Add caching for future user information, use caching
}
