package tokens_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
)

func (tr *TokensRepository) Create(ctx context.Context, m *models.Token) (int64, error) {
	query := `INSERT INTO $1 (jti, subject_id, type, revoked, issued_at, expires_at) VALUES ($2, $3, $4, $5, $6, $7) RETURNING id`

	var id int64
	err := tr.pool.QueryRow(ctx, query, tr.tableName, m.JTI, m.SubjectID, m.Type, m.Revoked, m.IssuedAt, m.ExpiresAt).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return -1, services.AlreadyExists
			}
		}

		return -1, services.InternalServer
	}

	return id, nil
}
