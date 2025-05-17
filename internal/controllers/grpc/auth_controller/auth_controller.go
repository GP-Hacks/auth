package auth_controller

import (
	"context"

	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services/credentials_service"
	"github.com/GP-Hacks/proto/pkg/api/auth"
	"github.com/GP-Hacks/proto/pkg/api/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (c *AuthController) SignUp(ctx context.Context, req *auth.SignUpRequest) (*emptypb.Empty, error) {
	var st models.UserStatus
	if req.User.Status == user.UserStatus_ADMIN {
		st = models.AdminUser
	} else if req.User.Status == user.UserStatus_DEFAULT {
		st = models.DefaultUser
	} else {
		return &emptypb.Empty{}, status.Errorf(codes.InvalidArgument, "invalid user status")
	}
	u := models.User{
		Email:       req.User.Email,
		FirstName:   req.User.FirstName,
		LastName:    req.User.LastName,
		Surname:     req.User.Surname,
		DateOfBirth: req.User.DateOfBirth.AsTime(),
		Status:      st,
	}
	return &emptypb.Empty{}, c.credentialsService.SignUp(ctx, req.Credentials.Email, req.Credentials.Password, &u)
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
