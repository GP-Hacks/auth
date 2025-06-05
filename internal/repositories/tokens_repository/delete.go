package tokens_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/utils/errs"
)

func (tr *TokensRepository) Delete(ctx context.Context, id int64) error {
	_, err := tr.pool.Exec(ctx,
		"DELETE FROM  issued_jwt_token WHERE id = $1", id)

	if err != nil {
		return errors.Join(errs.SomeError, err)
	}

	return nil
}
