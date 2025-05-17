package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/transaction"
)

type (
	IEmailTokensRepository interface {
		Save(ctx context.Context, token string, userID int64) error
		GetUserID(ctx context.Context, token string) (int64, error)
	}

	INotificationsAdapter interface {
		SendMail(m *models.Mail) error
	}

	ICredentialsRepository interface {
		GetById(ctx context.Context, id int64) (*models.Credentials, error)
		Create(ctx context.Context, m *models.Credentials) (int64, error)
		GetByEmail(ctx context.Context, email string) (*models.Credentials, error)
		Confirm(ctx context.Context, id int64) error
	}

	ITokensRepository interface {
		GetById(ctx context.Context, id int64) (*models.Token, error)
		Create(ctx context.Context, m *models.Token) (int64, error)
		Update(ctx context.Context, m *models.Token) error
		Delete(ctx context.Context, id int64) error
		GetByJTI(ctx context.Context, jti string) (*models.Token, error)
		RevokeAllWithSubjectId(ctx context.Context, subId int64) error
		RevokeByJTI(ctx context.Context, jti string) error
	}

	IUsersAdapter interface {
		Create(ctx context.Context, user *models.User) error
	}

	CredentialsService struct {
		emailTokensRepository IEmailTokensRepository
		credentialsRepository ICredentialsRepository
		tokensRepository      ITokensRepository
		userAdaper            IUsersAdapter
		notificationsAdapter  INotificationsAdapter
		txManager             transaction.TransactionManager
	}
)

func NewCredentialsService(er IEmailTokensRepository, cr ICredentialsRepository, tr ITokensRepository, ua IUsersAdapter, na INotificationsAdapter, tm transaction.TransactionManager) *CredentialsService {
	return &CredentialsService{
		emailTokensRepository: er,
		credentialsRepository: cr,
		tokensRepository:      tr,
		userAdaper:            ua,
		notificationsAdapter:  na,
		txManager:             tm,
	}
}
