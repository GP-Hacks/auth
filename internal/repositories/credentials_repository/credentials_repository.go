package credentials_repository

import "github.com/jackc/pgx/v5/pgxpool"

type CredentialsRepository struct {
	pool      *pgxpool.Pool
	tableName string
}

func NewCredentialsRepository(pool *pgxpool.Pool, tableName string) *CredentialsRepository {
	return &CredentialsRepository{
		pool:      pool,
		tableName: tableName,
	}
}
