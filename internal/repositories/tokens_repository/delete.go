package tokens_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/services"
	"github.com/rs/zerolog/log"
)

func (tr *TokensRepository) Delete(ctx context.Context, id int64) error {
	_, err := tr.pool.Exec(ctx,
		"DELETE FROM  issued_jwt_token WHERE id = $1", id)

	if err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServer
	}

	return nil
}
