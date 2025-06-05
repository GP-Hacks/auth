package email_tokens_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/go-redis/redis/v8"
)

func (r *EmailTokensRepository) GetUserID(ctx context.Context, token string) (int64, error) {
	val, err := r.client.Get(ctx, token).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, errs.NotFoundError
		}

		return 0, errors.Join(errs.SomeError, err)
	}
	var userID int64
	_, err = fmt.Sscan(val, &userID)
	if err != nil {
		return 0, errors.Join(errs.SomeError, err)
	}

	return userID, nil
}
