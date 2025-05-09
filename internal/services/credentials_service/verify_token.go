package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/golang-jwt/jwt/v5"
)

func (s *CredentialsService) verify_token(ctx context.Context, tokenString string, expectedType models.TokenType) (int64, string, error) {
	secretKey := "" // TODO: Replace to get real secret key

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return -1, "", services.InternalServer
	}

	if !token.Valid {
		return -1, "", services.InvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		typeToken, ok := claims["type"].(models.TokenType)
		if !ok {
			return -1, "", services.InvalidToken
		}
		if typeToken != expectedType {
			return -1, "", services.InvalidToken
		}

		jti, ok := claims["jti"].(string)
		if !ok {
			return -1, "", services.InvalidToken
		}

		subId, ok := claims["id"].(int64)
		if !ok {
			return -1, "", services.InvalidToken
		}

		t, err := s.tokensRepository.GetByJTI(ctx, jti)
		if err != nil {
			return -1, "", err
		}

		if t.Revoked || t.SubjectID != subId {
			return -1, "", services.InvalidToken
		}

		return subId, jti, nil
	}

	return -1, "", services.InvalidToken
}
