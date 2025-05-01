package credentials_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/jackc/pgx/v5"
)

func (cr *CredentialsRepository) GetById(ctx context.Context, id int64) (*models.Credentials, error) {
	query := `SELECT id, email, password FROM $1 WHERE id = $2`

	var credentials models.Credentials
	err := cr.pool.QueryRow(ctx, query, cr.tableName, id).Scan(
		&credentials.ID,
		&credentials.Email,
		&credentials.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		return nil, services.InternalServer
	}

	return &credentials, nil
}

func (cr *CredentialsRepository) GetByEmail(ctx context.Context, email string) (*models.Credentials, error) {
	query := `SELECT id, email, password FROM $1 WHERE email = $2`

	var credentials models.Credentials
	err := cr.pool.QueryRow(ctx, query, cr.tableName, email).Scan(
		&credentials.ID,
		&credentials.Email,
		&credentials.Password,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, services.NotFound
		}

		return nil, services.InternalServer
	}

	return &credentials, nil
}
