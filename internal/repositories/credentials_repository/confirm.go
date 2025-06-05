package credentials_repository

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/utils/errs"
)

func (r *CredentialsRepository) Confirm(ctx context.Context, id int64) error {
	q := `UPDATE users SET is_verification = true WHERE id = $1`

	if _, err := r.pool.Exec(ctx, q, id); err != nil {
		return errors.Join(errs.SomeError, err)
	}

	return nil
}
