package email_tokens_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/GP-Hacks/auth/internal/utils/errs"
)

func (r *EmailTokensRepository) Save(ctx context.Context, token string, userID int64) error {
	val := fmt.Sprint(userID)
	err := r.client.Set(ctx, token, val, config.Cfg.Redis.TokensTTL).Err()
	if err != nil {
		return errors.Join(errs.SomeError, err)
	}
	return nil
}
