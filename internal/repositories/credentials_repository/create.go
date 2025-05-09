package credentials_repository

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/rs/zerolog/log"
)

func (cr *CredentialsRepository) Create(ctx context.Context, m *models.Credentials) (int64, error) {
	query := `INSERT INTO credentials (email, password) VALUES ($1, $2) RETURNING id`

	var id int64
	err := cr.pool.QueryRow(ctx, query, m.Email, m.Password).Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				return -1, services.AlreadyExists
			}
		}

		log.Error().Msg(err.Error())
		return -1, services.InternalServer
	}

	return id, nil
}
