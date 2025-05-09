package tokens_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/services"
)

func (r *TokensRepository) RevokeByJTI(ctx context.Context, jti string) error {
	query := `UPDATE $1 SET revoked = true WHERE jti = $2`

	_, err := r.pool.Exec(ctx, query, r.tableName, jti)
	if err != nil {
		return services.InternalServer
	}

	return nil
}

func (r *TokensRepository) RevokeAllWithSubjectId(ctx context.Context, subId int64) error {
	query := `UPDATE $1 SET revoked = true WHERE subject_id = $2`

	_, err := r.pool.Exec(ctx, query, r.tableName, subId)
	if err != nil {
		return services.InternalServer
	}

	return nil
}
