package credentials_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
)

func (s *CredentialsService) Logout(ctx context.Context, access, refresh string) error {
	idFromAccess, jtiFromAccess, err := s.verifyToken(ctx, access, models.Access)
	if err != nil {
		if errors.Is(err, errs.UnauthorizedError) {
			return fmt.Errorf("Verify access token: %w", err)
		}
		return fmt.Errorf("Verify access token: %w: %v", errs.SomeError, err)
	}

	idFromRefresh, jtiFromRefresh, err := s.verifyToken(ctx, refresh, models.Refresh)
	if err != nil {
		if errors.Is(err, errs.UnauthorizedError) {
			return fmt.Errorf("Verify refresh token: %w", err)
		}
		return fmt.Errorf("Verify refresh token: %w: %v", errs.SomeError, err)
	}

	if idFromAccess != idFromRefresh {
		return fmt.Errorf("tokens id do not match: %w", errs.UnauthorizedError)
	}

	if err := s.tokensRepository.RevokeByJTI(ctx, jtiFromAccess); err != nil {
		if errors.Is(err, errs.NotFoundError) {
			return fmt.Errorf("Revoke access: %w: %v", errs.UnauthorizedError, err)
		}
		return fmt.Errorf("Revoke access: %w: %v", errs.SomeError, err)
	}

	if err := s.tokensRepository.RevokeByJTI(ctx, jtiFromRefresh); err != nil {
		if errors.Is(err, errs.NotFoundError) {
			return fmt.Errorf("Revoke refresh: %w: %v", errs.UnauthorizedError, err)
		}
		return fmt.Errorf("Revoke refresh: %w: %v", errs.SomeError, err)
	}

	return nil
}
