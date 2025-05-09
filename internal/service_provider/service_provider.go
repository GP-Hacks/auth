package service_provider

import (
	"github.com/GP-Hacks/auth/internal/controllers/grpc/auth_controller"
	"github.com/GP-Hacks/auth/internal/repositories/credentials_repository"
	"github.com/GP-Hacks/auth/internal/repositories/tokens_repository"
	"github.com/GP-Hacks/auth/internal/services/credentials_service"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceProvider struct {
	db *pgxpool.Pool

	tokensRepository      *tokens_repository.TokensRepository
	credentialsRepository *credentials_repository.CredentialsRepository

	credentialsService *credentials_service.CredentialsService

	authController *auth_controller.AuthController
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
