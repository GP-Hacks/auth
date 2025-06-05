package auth_controller

import (
	"context"
	"errors"

	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryDomainErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		var code codes.Code
		if errors.Is(err, errs.UnauthorizedError) {
			code = codes.Unauthenticated
		} else if errors.Is(err, errs.NotFoundError) {
			code = codes.NotFound
		} else if errors.Is(err, errs.AlreadyExistsError) {
			code = codes.AlreadyExists
		} else if errors.Is(err, errs.SomeError) {
			code = codes.Internal
		} else {
			code = codes.Unknown
		}

		log.Error().Msg(err.Error())
		return resp, status.Error(code, "")
	}
}
