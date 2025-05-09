package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/GP-Hacks/auth/internal/utils/hasher"
)

func (s *CredentialsService) SignUp(ctx context.Context, email, password string) error {
	// TODO: Add validate password

	hash, err := hasher.HashPassword(password)
	if err != nil {
		return services.InternalServer
	}

	cred := models.Credentials{
		Email:    email,
		Password: hash,
	}

	if _, err := s.credentialsRepository.Create(ctx, &cred); err != nil {
		return err
	}

	return nil
}
