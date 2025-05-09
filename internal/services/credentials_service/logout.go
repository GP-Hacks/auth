package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
)

func (s *CredentialsService) Logout(ctx context.Context, access, refresh string) error {
	idFromAccess, jtiFromAccess, err := s.verify_token(ctx, access, models.Access)
	if err != nil {
		return err
	}

	idFromRefresh, jtiFromRefresh, err := s.verify_token(ctx, refresh, models.Refresh)
	if err != nil {
		return err
	}

	if idFromAccess != idFromRefresh {
		return services.InvalidToken
	}

	if err := s.tokensRepository.RevokeByJTI(ctx, jtiFromAccess); err != nil {
		return err
	}

	if err := s.tokensRepository.RevokeByJTI(ctx, jtiFromRefresh); err != nil {
		return err
	}

	return nil
}
