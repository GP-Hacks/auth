package credentials_service

import (
	"context"
	"time"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/GP-Hacks/auth/internal/utils/hasher"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (s *CredentialsService) SignIn(ctx context.Context, email, password string) (string, string, error) {
	cred, err := s.credentialsRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if !cred.IsVerification {
		return "", "", services.NotVerification
	}

	if !hasher.ValidatePassword(password, cred.Password) {
		return "", "", services.InvalidCredentials
	}

	access, accessString, err := s.createJWTToken(cred.ID, email, models.Access)
	if err != nil {
		return "", "", err
	}

	refresh, refreshString, err := s.createJWTToken(cred.ID, email, models.Refresh)
	if err != nil {
		return "", "", err
	}

	if _, err := s.tokensRepository.Create(ctx, access); err != nil {
		return "", "", err
	}
	if _, err := s.tokensRepository.Create(ctx, refresh); err != nil {
		return "", "", err
	}

	return accessString, refreshString, nil
}

func (s *CredentialsService) createJWTToken(credentialsId int64, email string, typeToken models.TokenType) (*models.Token, string, error) {
	var (
		lifeTime int64
		err      error
		exp      int64
	)

	if typeToken == models.Refresh {
		// lifeTime, err = strconv.ParseInt(os.Getenv(refreshLifeTimeName), 0, 64)
		lifeTime = 100
		if err != nil || lifeTime < 1 {
			return nil, "", services.InternalServer
		}

		exp = time.Now().Add(time.Hour * 24 * time.Duration(lifeTime)).Unix()
	} else if typeToken == models.Access {
		// lifeTime, err = strconv.ParseInt(os.Getenv(accessLifeTimeName), 0, 64)
		lifeTime = 100
		if err != nil || lifeTime < 1 {
			return nil, "", services.InternalServer
		}

		exp = time.Now().Add(time.Minute * time.Duration(lifeTime)).Unix()
	} else {
		return nil, "", services.InternalServer
	}

	jti := uuid.New().String()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    credentialsId,
			"email": email,
			"exp":   exp,
			"jti":   jti,
			"type":  typeToken,
		},
	)

	tokenString, err := s.getTokenString(token)
	if err != nil {
		return nil, "", err
	}

	return &models.Token{
		JTI:       jti,
		SubjectID: credentialsId,
		Type:      typeToken,
		Revoked:   false,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Unix(exp, 0),
	}, tokenString, nil
}

func (s *CredentialsService) getTokenString(token *jwt.Token) (string, error) {
	// secretKey := os.Getenv(secretKeyName)
	// if len(secretKey) == 0 {
	// return "", utils.InternalServerError
	// }
	secretKey := ""

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", services.InternalServer
	}

	return tokenString, nil
}
