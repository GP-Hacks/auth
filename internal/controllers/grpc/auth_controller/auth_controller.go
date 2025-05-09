package auth_controller

import (
	"context"

	"github.com/GP-Hacks/auth/internal/services/credentials_service"
	"github.com/GP-Hacks/proto/pkg/api/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	auth.UnimplementedAuthServiceServer
	credentialsService *credentials_service.CredentialsService
}

func (c *AuthController) SignUp(ctx context.Context, req *auth.SignUpRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, c.credentialsService.SignUp(ctx, req.Credentials.Email, req.Credentials.Password)
}

func (c *AuthController) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	access, refresh, err := c.credentialsService.SignIn(ctx, req.Credentials.Email, req.Credentials.Password)
	if err != nil {
		return nil, err
	}

	return &auth.SignInResponse{
		Tokens: &auth.Tokens{
			Access:  access,
			Refresh: refresh,
		},
	}, nil
}

func (c *AuthController) VerifyAccessToken(ctx context.Context, req *auth.VerifyAccessTokenRequest) (*auth.VerifyAccessTokenResponse, error) {
	id, err := c.credentialsService.VerifyAccessToken(ctx, req.Access)
	if err != nil {
		return nil, err
	}

	return &auth.VerifyAccessTokenResponse{
		UserId: id,
	}, nil
}

func (c *AuthController) RefreshTokens(ctx context.Context, req *auth.RefreshTokensRequest) (*auth.RefreshTokensResponse, error) {
	access, refresh, err := c.credentialsService.RefreshTokens(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokensResponse{
		Tokens: &auth.Tokens{
			Access:  access,
			Refresh: refresh,
		},
	}, nil
}

func (c *AuthController) Logout(ctx context.Context, req *auth.LogoutRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, c.credentialsService.Logout(ctx, req.Tokens.Access, req.Tokens.Refresh)
}
