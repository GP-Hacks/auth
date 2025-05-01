package tokens_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/services"
)

func (tr *TokensRepository) Delete(ctx context.Context, id int64) error {
	_, err := tr.pool.Exec(ctx,
		"DELETE FROM $1 WHERE id = $2", tr.tableName, id)

	if err != nil {
		return services.InternalServer
	}

	return nil
}
