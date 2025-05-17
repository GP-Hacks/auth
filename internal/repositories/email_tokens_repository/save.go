package email_tokens_repository

import (
	"context"
	"fmt"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/rs/zerolog/log"
)

func (r *EmailTokensRepository) Save(ctx context.Context, token string, userID int64) error {
	val := fmt.Sprint(userID)
	err := r.client.Set(ctx, token, val, config.Cfg.Redis.TokensTTL).Err()
	if err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServer
	}
	return nil
}
