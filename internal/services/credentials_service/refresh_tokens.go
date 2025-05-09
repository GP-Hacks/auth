package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
)

func (s *CredentialsService) RefreshTokens(ctx context.Context, refresh string) (string, string, error) {
	idFromRefresh, _, err := s.verify_token(ctx, refresh, models.Refresh)
	if err != nil {
		return "", "", err
	}

	if err := s.tokensRepository.RevokeAllWithSubjectId(ctx, idFromRefresh); err != nil {
		return "", "", err
	}

	cred, err := s.credentialsRepository.GetById(ctx, idFromRefresh)
	if err != nil {
		return "", "", err
	}

	newRefresh, refreshString, err := s.createJWTToken(cred.ID, cred.Email, models.Refresh)
	if err != nil {
		return "", "", err
	}

	newAccess, accessString, err := s.createJWTToken(cred.ID, cred.Email, models.Access)
	if err != nil {
		return "", "", err
	}

	if _, err := s.tokensRepository.Create(ctx, newRefresh); err != nil {
		return "", "", err
	}

	if _, err := s.tokensRepository.Create(ctx, newAccess); err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}
