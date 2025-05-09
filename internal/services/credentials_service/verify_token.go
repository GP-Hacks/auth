package credentials_service

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
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
		log.Debug().Msg("Token is not valid")
		return -1, "", services.InvalidToken
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key := range claims {
			log.Debug().Msg(key)
		}
		typeFloat, ok := claims["type"].(float64)
		if !ok {
			log.Debug().Msg("Token is not type")
			return -1, "", services.InvalidToken
		}
		typeToken := models.TokenType(int(typeFloat))
		if typeToken != expectedType {
			log.Debug().Msg("Token is type not expected")
			return -1, "", services.InvalidToken
		}

		jti, ok := claims["jti"].(string)
		if !ok {
			log.Debug().Msg("Token is not jti")
			return -1, "", services.InvalidToken
		}

		subIdF, ok := claims["id"].(float64)
		if !ok {
			log.Debug().Msg("Token is not subId")
			return -1, "", services.InvalidToken
		}
		subId := int64(subIdF)

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
