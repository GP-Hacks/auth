package credentials_service

import (
	"context"

	"github.com/rs/zerolog/log"
)

func (s *CredentialsService) ResendConfirmationEmail(ctx context.Context, email string) {
	cred, err := s.credentialsRepository.GetByEmail(ctx, email)
	if err != nil {
		log.Error().Msg(err.Error())
	}

	s.sendConfirmationEmail(cred)
}
