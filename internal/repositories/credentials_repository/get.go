package credentials_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/jackc/pgx/v5"
)

func (cr *CredentialsRepository) GetById(ctx context.Context, id int64) (*models.Credentials, error) {
	query := `SELECT id, email, password, is_verification FROM credentials WHERE id = $1`

	var credentials models.Credentials
	err := cr.pool.QueryRow(ctx, query, id).Scan(
		&credentials.ID,
		&credentials.Email,
		&credentials.Password,
		&credentials.IsVerification,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFoundError
		}

		return nil, errors.Join(errs.SomeError, err)
	}

	return &credentials, nil
}

func (cr *CredentialsRepository) GetByEmail(ctx context.Context, email string) (*models.Credentials, error) {
	query := `SELECT id, email, password, is_verification FROM credentials WHERE email = $1`

	var credentials models.Credentials
	err := cr.pool.QueryRow(ctx, query, email).Scan(
		&credentials.ID,
		&credentials.Email,
		&credentials.Password,
		&credentials.IsVerification,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errs.NotFoundError
		}

		return nil, errors.Join(errs.SomeError, err)
	}

	return &credentials, nil
}
