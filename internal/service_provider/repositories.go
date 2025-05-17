package service_provider

import (
	"github.com/GP-Hacks/auth/internal/repositories/credentials_repository"
	"github.com/GP-Hacks/auth/internal/repositories/email_tokens_repository"
	"github.com/GP-Hacks/auth/internal/repositories/tokens_repository"
)

func (s *ServiceProvider) TokensRepository() *tokens_repository.TokensRepository {
	if s.tokensRepository == nil {
		s.tokensRepository = tokens_repository.NewTokensRepository("issued_jwt_token", s.DB())
	}

	return s.tokensRepository
}

func (s *ServiceProvider) CredentialsRepository() *credentials_repository.CredentialsRepository {
	if s.credentialsRepository == nil {
		s.credentialsRepository = credentials_repository.NewCredentialsRepository(s.DB(), "credentials")
	}

	return s.credentialsRepository
}

func (s *ServiceProvider) EmailTokensRepository() *email_tokens_repository.EmailTokensRepository {
	if s.emailTokensRepository == nil {
		s.emailTokensRepository = email_tokens_repository.NewEmailTokensRepository(s.RedisClient())
	}

	return s.emailTokensRepository
}
