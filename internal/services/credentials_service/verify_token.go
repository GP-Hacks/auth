package credentials_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/golang-jwt/jwt/v5"
)

func (s *CredentialsService) verifyToken(ctx context.Context, tokenString string, expectedType models.TokenType) (int64, string, error) {
	secretKey := config.Cfg.JWT.SecretKey

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, "", fmt.Errorf("parse token: %w: %v", errs.UnauthorizedError, err)
	}

	if !token.Valid {
		return 0, "", fmt.Errorf("invalid token: %w", errs.UnauthorizedError)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		typeFloat, ok := claims["type"].(float64)
		if !ok {
			return -1, "", fmt.Errorf("invalid token: %w: none type", errs.UnauthorizedError)
		}
		typeToken := models.TokenType(int(typeFloat))
		if typeToken != expectedType {
			return -1, "", fmt.Errorf("invalid token: %w: not expected type", errs.UnauthorizedError)
		}

		jti, ok := claims["jti"].(string)
		if !ok {
			return -1, "", fmt.Errorf("invalid token: %w: none jti", errs.UnauthorizedError)
		}

		subIdF, ok := claims["id"].(float64)
		if !ok {
			return -1, "", fmt.Errorf("invalid token: %w: none sub id", errs.UnauthorizedError)
		}
		subId := int64(subIdF)

		t, err := s.tokensRepository.GetByJTI(ctx, jti)
		if err != nil {
			if errors.Is(err, errs.NotFoundError) {
				return -1, "", fmt.Errorf("get by jti: %w: %v", errs.UnauthorizedError, err)
			}
		}

		if t.Revoked || t.SubjectID != subId {
			return -1, "", fmt.Errorf("invalid token: %w", errs.UnauthorizedError)
		}

		return subId, jti, nil
	}

	return -1, "", fmt.Errorf("invalid token: %w", errs.UnauthorizedError)
}
