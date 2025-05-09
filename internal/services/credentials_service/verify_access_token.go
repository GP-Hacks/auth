package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
)

func (s *CredentialsService) VerifyAccessToken(ctx context.Context, token string) (int64, error) {
	id, _, err := s.verify_token(ctx, token, models.Access)
	if err != nil {
		return -1, err
	}

	return id, nil
}
