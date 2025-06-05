package tokens_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/jackc/pgx/v5"
)

func (tr *TokensRepository) GetById(ctx context.Context, id int64) (*models.Token, error) {
	query := `SELECT id, jti, subject_id, token_type, revoked, issued_at, expires_at FROM issued_jwt_token WHERE id = $1`

	var token models.Token
	err := tr.pool.QueryRow(ctx, query, id).Scan(
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
			return nil, errs.NotFoundError
		}

		return nil, errors.Join(errs.SomeError, err)
	}

	return &token, nil
}

func (tr *TokensRepository) GetByJTI(ctx context.Context, jti string) (*models.Token, error) {
	query := `SELECT id, jti, subject_id, token_type, revoked, issued_at, expires_at FROM issued_jwt_token WHERE jti = $1`

	var token models.Token
	err := tr.pool.QueryRow(ctx, query, jti).Scan(
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
			return nil, errs.NotFoundError
		}

		return nil, errors.Join(errs.SomeError, err)
	}

	return &token, nil
}
