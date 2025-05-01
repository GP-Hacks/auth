package credentials_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
)

func (cr *CredentialsRepository) Create(ctx context.Context, m *models.Credentials) (int64, error) {
	query := `INSERT INTO $1 (email, password) VALUES ($2, $3) RETURNING id`

	var id int64
	err := cr.pool.QueryRow(ctx, query, cr.tableName, m.Email, m.Password).Scan(&id)
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
