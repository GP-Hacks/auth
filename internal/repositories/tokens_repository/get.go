package tokens_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/jackc/pgx/v5"
)

func (tr *TokensRepository) GetByID(ctx context.Context, id int64) (*models.Token, error) {
	query := `SELECT id, jti, subject_id, type, revoked, issued_at, expires_at FROM $1 WHERE id = $2`

	var token models.Token
	err := tr.pool.QueryRow(ctx, query, tr.tableName, id).Scan(
		&token.ID,
		&token.JTI,
		&token.SubjectID,
		&token.Type,
		&token.Revoked,
		&token.IssuedAt,
		&token.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		return nil, services.InternalServer
	}

	return &token, nil
}

func (tr *TokensRepository) GetByJTI(ctx context.Context, jti string) (*models.Token, error) {
	query := `SELECT id, jti, subject_id, type, revoked, issued_at, expires_at FROM $1 WHERE jti = $2`

	var token models.Token
	err := tr.pool.QueryRow(ctx, query, tr.tableName, jti).Scan(
		&token.ID,
		&token.JTI,
		&token.SubjectID,
		&token.Type,
		&token.Revoked,
		&token.IssuedAt,
		&token.ExpiresAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		return nil, services.InternalServer
	}

	return &token, nil
}
