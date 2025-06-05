package credentials_service

import (
	"context"
	"fmt"
)

func (s *CredentialsService) ConfirmationEmail(ctx context.Context, token string) error {
	id, err := s.emailTokensRepository.GetUserID(ctx, token)
	if err != nil {
		return fmt.Errorf("Get user id by token: %w", err)
	}

	if err := s.credentialsRepository.Confirm(ctx, id); err != nil {
		return fmt.Errorf("Confirm account: %w", err)
	}

	return nil
}
