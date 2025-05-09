package tokens_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
)

func (tr *TokensRepository) Update(ctx context.Context, m *models.Token) error {
	query := `UPDATE $1 SET
		jti = $2,
		subject_id = $3,
		type = $4,
		revoked = $5,
		issued_at = $6,
		expires_at = $7
		WHERE id = $8
	`

	_, err := tr.pool.Exec(ctx, query,
		tr.tableName,
		m.JTI,
		m.SubjectID,
		m.Type,
		m.Revoked,
		m.IssuedAt,
		m.ExpiresAt,
		m.ID,
	)
	if err != nil {
		return services.InternalServer
	}
	return nil

}
