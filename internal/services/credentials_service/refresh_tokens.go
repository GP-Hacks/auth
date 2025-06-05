package credentials_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
)

func (s *CredentialsService) RefreshTokens(ctx context.Context, refresh string) (string, string, error) {
	idFromRefresh, jti, err := s.verifyToken(ctx, refresh, models.Refresh)
	if err != nil {
		if errors.Is(err, errs.UnauthorizedError) {
			return "", "", fmt.Errorf("Verify refresh token: %w", err)
		}
		return "", "", fmt.Errorf("Verify refresh token: %w: %v", errs.SomeError, err)
	}

	if err := s.tokensRepository.RevokeByJTI(ctx, jti); err != nil {
		if errors.Is(err, errs.NotFoundError) {
			return "", "", fmt.Errorf("Revoke refresh: %w: %v", errs.UnauthorizedError, err)
		}
		return "", "", fmt.Errorf("Revoke refresh: %w: %v", errs.SomeError, err)
	}

	cred, err := s.credentialsRepository.GetById(ctx, idFromRefresh)
	if err != nil {
		if errors.Is(err, errs.NotFoundError) {
			return "", "", fmt.Errorf("Get credentials: %w: %v", errs.UnauthorizedError, err)
		}
		return "", "", fmt.Errorf("Get credentials: %w: %v", errs.SomeError, err)
	}

	newRefresh, refreshString, err := s.createJWTToken(cred.ID, cred.Email, models.Refresh)
	if err != nil {
		return "", "", fmt.Errorf("Create refresh: %w", err)
	}

	newAccess, accessString, err := s.createJWTToken(cred.ID, cred.Email, models.Access)
	if err != nil {
		return "", "", fmt.Errorf("Create access: %w", err)
	}

	if _, err := s.tokensRepository.Create(ctx, newRefresh); err != nil {
		return "", "", fmt.Errorf("save refresh: %w", err)
	}

	if _, err := s.tokensRepository.Create(ctx, newAccess); err != nil {
		return "", "", fmt.Errorf("save refresh: %w", err)
	}

	return accessString, refreshString, nil
}
