package credentials_service

import "context"

func (s *CredentialsService) ConfirmationEmail(ctx context.Context, token string) error {
	id, err := s.emailTokensRepository.GetUserID(ctx, token)
	if err != nil {
		return err
	}

	if err := s.credentialsRepository.Confirm(ctx, id); err != nil {
		return err
	}

	return nil
}
