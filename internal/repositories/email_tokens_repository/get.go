package email_tokens_repository

import (
	"context"
	"fmt"

	"github.com/GP-Hacks/auth/internal/services"
	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog/log"
)

func (r *EmailTokensRepository) GetUserID(ctx context.Context, token string) (int64, error) {
	val, err := r.client.Get(ctx, token).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, services.NotFound
		}
		log.Error().Msg(err.Error())
		return 0, services.InternalServer
	}
	var userID int64
	_, err = fmt.Sscan(val, &userID)
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, services.InternalServer
	}

	return userID, nil
}
