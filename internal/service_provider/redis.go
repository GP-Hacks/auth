package service_provider

import (
	"github.com/GP-Hacks/auth/internal/config"
	"github.com/go-redis/redis/v8"
)

func (s *ServiceProvider) RedisClient() *redis.Client {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(&redis.Options{
			Addr:     config.Cfg.Redis.Address,
			Password: config.Cfg.Redis.Password,
			DB:       config.Cfg.Redis.DB,
		})
	}

	return s.redisClient
}
