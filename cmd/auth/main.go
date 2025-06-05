package main

import (
	"net"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/GP-Hacks/auth/internal/controllers/grpc/auth_controller"
	"github.com/GP-Hacks/auth/internal/service_provider"
	"github.com/GP-Hacks/auth/internal/utils/logger"
	proto "github.com/GP-Hacks/proto/pkg/api/auth"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config.LoadConfig("./config")
	logger.SetupLogger()
	serviceProvider := service_provider.NewServiceProvider()

	log.Info().Msg("Init app")

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(auth_controller.UnaryDomainErrorInterceptor()))
	reflection.Register(grpcServer)

	proto.RegisterAuthServiceServer(grpcServer, serviceProvider.AuthController())

	list, err := net.Listen("tcp", ":"+config.Cfg.Grpc.Port)
	if err != nil {
		log.Fatal().Msg("Failed start listen port")
	}

	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal().Msg("Failed serve grpc")
	}
}
