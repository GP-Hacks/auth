package credentials_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/GP-Hacks/auth/internal/utils/transaction"
	"github.com/jackc/pgx/v5/pgconn"
)

func (cr *CredentialsRepository) Create(ctx context.Context, m *models.Credentials) (int64, error) {
	tx, err := transaction.TxFromCtx(ctx)
	query := `INSERT INTO credentials (email, password, is_verification) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	if err != nil {
		err = cr.pool.QueryRow(ctx, query, m.Email, m.Password, m.IsVerification).Scan(&id)
	} else {
		err = tx.QueryRow(ctx, query, m.Email, m.Password, m.IsVerification).Scan(&id)
	}
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return -1, errs.SomeError
			}
		}

		return -1, errors.Join(errs.SomeError, err)
	}

	return id, nil
}
