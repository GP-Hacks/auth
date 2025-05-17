package service_provider

import (
	"github.com/GP-Hacks/auth/internal/adapters/notifications_adapter"
	"github.com/GP-Hacks/auth/internal/adapters/users_adapter"
	"github.com/GP-Hacks/auth/internal/controllers/grpc/auth_controller"
	"github.com/GP-Hacks/auth/internal/repositories/credentials_repository"
	"github.com/GP-Hacks/auth/internal/repositories/email_tokens_repository"
	"github.com/GP-Hacks/auth/internal/repositories/tokens_repository"
	"github.com/GP-Hacks/auth/internal/services/credentials_service"
	desc "github.com/GP-Hacks/proto/pkg/api/user"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

type ServiceProvider struct {
	db              *pgxpool.Pool
	usersClient     desc.UserServiceClient
	usersConnection *grpc.ClientConn
	rabbitMqCh      *amqp091.Channel
	rabbitmqConn    *amqp091.Connection
	redisClient     *redis.Client

	emailTokensRepository *email_tokens_repository.EmailTokensRepository
	usersAdapter          *users_adapter.UsersAdapter
	notificationsAdapter  *notifications_adapter.NotificationsAdapter
	tokensRepository      *tokens_repository.TokensRepository
	credentialsRepository *credentials_repository.CredentialsRepository

	credentialsService *credentials_service.CredentialsService

	authController *auth_controller.AuthController
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
