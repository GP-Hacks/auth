package credentials_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/services"
	"github.com/rs/zerolog/log"
)

func (r *CredentialsRepository) Confirm(ctx context.Context, id int64) error {
	q := `UPDATE users SET is_verification = true WHERE id = $1`

	if _, err := r.pool.Exec(ctx, q, id); err != nil {
		log.Error().Msg(err.Error())

		return services.InternalServer
	}

	return nil
}
