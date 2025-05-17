package service_provider

import (
	"github.com/GP-Hacks/auth/internal/adapters/notifications_adapter"
	"github.com/GP-Hacks/auth/internal/adapters/users_adapter"
)

func (s *ServiceProvider) UsersAdapter() *users_adapter.UsersAdapter {
	if s.usersAdapter == nil {
		s.usersAdapter = users_adapter.NewUsersAdapter(s.UsersClient())
	}

	return s.usersAdapter
}

func (s *ServiceProvider) NotificationsAdapter() *notifications_adapter.NotificationsAdapter {
	if s.notificationsAdapter == nil {
		s.notificationsAdapter = notifications_adapter.NewNotificationsAdapter(s.RabbitMQChannel())
	}

	return s.notificationsAdapter
}
