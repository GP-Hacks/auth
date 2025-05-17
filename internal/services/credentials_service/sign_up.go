package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/GP-Hacks/auth/internal/utils/hasher"
	"github.com/rs/zerolog/log"
)

func (s *CredentialsService) SignUp(ctx context.Context, email, password string, u *models.User) error {
	// TODO: Add validate password

	hash, err := hasher.HashPassword(password)
	if err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServer
	}

	cred := models.Credentials{
		Email:          email,
		Password:       hash,
		IsVerification: false,
	}

	return s.txManager.Do(ctx, func(txCtx context.Context) error {
		id, err := s.credentialsRepository.Create(txCtx, &cred)
		if err != nil {
			return err
		}

		u.ID = id
		if err := s.userAdaper.Create(ctx, u); err != nil {
			return err
		}

		go s.sendConfirmationEmail(u)

		return nil
	})
}
