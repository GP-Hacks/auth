package tokens_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/jackc/pgx/v5/pgconn"
)

func (tr *TokensRepository) Create(ctx context.Context, m *models.Token) (int64, error) {
	query := `INSERT INTO issued_jwt_token (jti, subject_id, token_type, revoked, issued_at, expires_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int64
	err := tr.pool.QueryRow(ctx, query, m.JTI, m.SubjectID, m.Type, m.Revoked, m.IssuedAt, m.ExpiresAt).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return -1, errs.AlreadyExistsError
			}
		}

		return -1, errors.Join(errs.SomeError, err)
	}

	return id, nil
}
