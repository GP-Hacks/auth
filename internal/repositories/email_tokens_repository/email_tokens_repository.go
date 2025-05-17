package email_tokens_repository

import "github.com/go-redis/redis/v8"

type EmailTokensRepository struct {
	client *redis.Client
}

func NewEmailTokensRepository(cl *redis.Client) *EmailTokensRepository {
	return &EmailTokensRepository{
		client: cl,
	}
}
