package users_adapter

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	desc "github.com/GP-Hacks/proto/pkg/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *UsersAdapter) Create(ctx context.Context, user *models.User) error {
	proto := desc.CreateRequest{
		Id: user.ID,
		User: &desc.User{
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Surname:     user.Surname,
			DateOfBirth: timestamppb.New(user.DateOfBirth),
		},
	}

	_, err := a.client.Create(ctx, &proto)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.AlreadyExists:
				return errs.AlreadyExistsError
			default:
				return errs.SomeError
			}
		}
		return errs.SomeError
	}

	return nil
}
