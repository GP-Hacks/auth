package users_adapter

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	desc "github.com/GP-Hacks/proto/pkg/api/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *UsersAdapter) Create(ctx context.Context, user *models.User) error {
	var st desc.UserStatus
	if user.Status == models.DefaultUser {
		st = desc.UserStatus_DEFAULT
	} else if user.Status == models.AdminUser {
		st = desc.UserStatus_ADMIN
	} else {
		log.Error().Msg("Invalid user status")
		return services.InternalServer
	}

	proto := desc.CreateRequest{
		User: &desc.User{
			Id:          user.ID,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Surname:     user.Surname,
			DateOfBirth: timestamppb.New(user.DateOfBirth),
			Status:      st,
		},
	}

	_, err := a.client.Create(ctx, &proto)
	if err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServer
	}

	return nil
}
