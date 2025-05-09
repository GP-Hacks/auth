package service_provider

import "github.com/GP-Hacks/auth/internal/controllers/grpc/auth_controller"

func (s *ServiceProvider) AuthController() *auth_controller.AuthController {
	if s.authController == nil {
		s.authController = auth_controller.NewAuthController(s.CredentialsService())
	}

	return s.authController
}
