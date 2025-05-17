package service_provider

import (
	desc "github.com/GP-Hacks/proto/pkg/api/user"
)

func (s *ServiceProvider) UsersClient() desc.UserServiceClient {
	if s.usersClient == nil {
		s.usersClient = desc.NewUserServiceClient(s.UsersConnection())
	}

	return s.usersClient
}
