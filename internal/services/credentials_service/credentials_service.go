package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
)

type (
	ICredentialsRepository interface {
		GetById(ctx context.Context, id int64) (*models.Credentials, error)
		Create(ctx context.Context, m *models.Credentials) (int64, error)
		GetByEmail(ctx context.Context, email string) (*models.Credentials, error)
	}

	ITokensRepository interface {
		GetById(ctx context.Context, id int64) (*models.Token, error)
		Create(ctx context.Context, m *models.Token) (int64, error)
		Update(ctx context.Context, m *models.Token) error
		Delete(ctx context.Context, id int64) error
		GetByJTI(ctx context.Context, jti string) (*models.Token, error)
	}
)
