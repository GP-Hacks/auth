package service_provider

import "github.com/GP-Hacks/auth/internal/services/credentials_service"

func (s *ServiceProvider) CredentialsService() *credentials_service.CredentialsService {
	if s.credentialsService == nil {
		s.credentialsService = credentials_service.NewCredentialsService(s.EmailTokensRepository(), s.CredentialsRepository(), s.TokensRepository(), s.UsersAdapter(), s.NotificationsAdapter(), s.TxManager())
	}

	return s.credentialsService
}
