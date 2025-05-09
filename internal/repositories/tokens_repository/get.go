package tokens_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
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
			return nil, services.NotFound
		}

		log.Error().Msg(err.Error())
		return nil, services.InternalServer
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
			return nil, services.NotFound
		}

		log.Error().Msg(err.Error())
		return nil, services.InternalServer
	}

	return &token, nil
}
