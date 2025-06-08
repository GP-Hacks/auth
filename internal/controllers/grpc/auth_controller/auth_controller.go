package auth_controller

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services/credentials_service"
	"github.com/GP-Hacks/proto/pkg/api/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthController struct {
	auth.UnimplementedAuthServiceServer
	credentialsService *credentials_service.CredentialsService
}

func NewAuthController(cs *credentials_service.CredentialsService) *AuthController {
	return &AuthController{
		credentialsService: cs,
	}
}

func (c *AuthController) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*emptypb.Empty, error) {
	if err := c.credentialsService.ConfirmationEmail(ctx, req.Token); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (c *AuthController) SignUp(ctx context.Context, req *auth.SignUpRequest) (*emptypb.Empty, error) {
	u := models.User{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Surname:     req.Surname,
		DateOfBirth: req.DateOfBirth.AsTime(),
	}
	return &emptypb.Empty{}, c.credentialsService.SignUp(ctx, req.Email, req.Password, &u)
}

func (c *AuthController) SignIn(ctx context.Context, req *auth.SignInRequest) (*auth.SignInResponse, error) {
	access, refresh, err := c.credentialsService.SignIn(ctx, req.Email, req.Password)
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

func (c *AuthController) ResendConfirmationMail(ctx context.Context, req *auth.ResendConfirmationMailRequest) (*emptypb.Empty, error) {
	c.credentialsService.ResendConfirmationEmail(ctx, req.Email)
	return &emptypb.Empty{}, nil
}
