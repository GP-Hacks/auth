package tokens_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/utils/errs"
)

func (r *TokensRepository) RevokeByJTI(ctx context.Context, jti string) error {
	query := `UPDATE issued_jwt_token SET revoked = true WHERE jti = $1`

	_, err := r.pool.Exec(ctx, query, jti)
	if err != nil {
		return errors.Join(errs.SomeError, err)
	}

	return nil
}

func (r *TokensRepository) RevokeAllWithSubjectId(ctx context.Context, subId int64) error {
	query := `UPDATE issued_jwt_token SET revoked = true WHERE subject_id = $1`

	_, err := r.pool.Exec(ctx, query, subId)
	if err != nil {
		return errors.Join(errs.SomeError, err)
	}

	return nil
}
