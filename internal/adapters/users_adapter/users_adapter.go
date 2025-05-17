package users_adapter

import (
	desc "github.com/GP-Hacks/proto/pkg/api/user"
)

type UsersAdapter struct {
	desc.UserServiceClient
	client desc.UserServiceClient
}

func NewUsersAdapter(client desc.UserServiceClient) *UsersAdapter {
	return &UsersAdapter{client: client}
}
