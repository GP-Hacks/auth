package service_provider

import (
	"github.com/GP-Hacks/auth/internal/config"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ServiceProvider) UsersConnection() *grpc.ClientConn {
	if s.usersConnection == nil {
		conn, err := grpc.Dial(
			config.Cfg.Grpc.UsersServiceAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		s.usersConnection = conn
	}

	return s.usersConnection
}
