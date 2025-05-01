package tokens_repository

import "github.com/jackc/pgx/v5/pgxpool"

type TokensRepository struct {
	tableName string
	pool      *pgxpool.Pool
}

func NewTokensRepository(tableName string, pool *pgxpool.Pool) *TokensRepository {
	return &TokensRepository{
		tableName: tableName,
		pool:      pool,
	}
}
