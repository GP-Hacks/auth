package credentials_service

import (
	"context"
	"errors"
	"fmt"
	"unicode"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/GP-Hacks/auth/internal/utils/hasher"
)

func (s *CredentialsService) SignUp(ctx context.Context, email, password string, u *models.User) error {
	if err := s.validatePassword(password); err != nil {
		return err
	}

	hash, err := hasher.HashPassword(password)
	if err != nil {
		return fmt.Errorf("hash password: %w: %v", errs.SomeError, err)
	}

	cred := models.Credentials{
		Email:          email,
		Password:       hash,
		IsVerification: false,
	}

	return s.txManager.Do(ctx, func(txCtx context.Context) error {
		id, err := s.credentialsRepository.Create(txCtx, &cred)
		if err != nil {
			if errors.Is(err, errs.AlreadyExistsError) {
				return fmt.Errorf("create user: %w", err)
			}
			return fmt.Errorf("create user: %w: %v", errs.SomeError, err)
		}

		u.ID = id
		cred.ID = id
		if err := s.userAdaper.Create(ctx, u); err != nil {
			if errors.Is(err, errs.AlreadyExistsError) {
				return fmt.Errorf("create user: %w", err)
			}
			return fmt.Errorf("create user: %w: %v", errs.SomeError, err)
		}

		go s.sendConfirmationEmail(&cred)

		return nil
	})
}

func (s *CredentialsService) validatePassword(pw string) error {
	var (
		hasMinLen  = false
		hasLower   = false
		hasUpper   = false
		hasDigit   = false
		hasSpecial = false
	)

	if len(pw) >= 8 {
		hasMinLen = true
	}

	for _, r := range pw {
		switch {
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsDigit(r):
			hasDigit = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}
	}

	if !hasMinLen {
		return fmt.Errorf("not has min len")
	}
	if !hasLower {
		return fmt.Errorf("not has lower")
	}
	if !hasUpper {
		return fmt.Errorf("not has upper")
	}
	if !hasDigit {
		return fmt.Errorf("not has digit")
	}
	if !hasSpecial {
		return fmt.Errorf("not has special symbol")
	}
	return nil
}
