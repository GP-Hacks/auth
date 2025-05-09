package tokens_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/rs/zerolog/log"
)

func (tr *TokensRepository) Update(ctx context.Context, m *models.Token) error {
	query := `UPDATE issued_jwt_token SET
		jti = $1,
		subject_id = $2,
		token_type = $3,
		revoked = $4,
		issued_at = $5,
		expires_at = $6
		WHERE id = $7
	`

	_, err := tr.pool.Exec(ctx, query,
		m.JTI,
		m.SubjectID,
		m.Type,
		m.Revoked,
		m.IssuedAt,
		m.ExpiresAt,
		m.ID,
	)
	if err != nil {
		log.Error().Msg(err.Error())
		return services.InternalServer
	}
	return nil

}
