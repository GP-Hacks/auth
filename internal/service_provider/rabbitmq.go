package service_provider

import (
	"github.com/GP-Hacks/auth/internal/config"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func (s *ServiceProvider) RabbitMQConnection() *amqp091.Connection {
	if s.rabbitmqConn == nil {
		conn, err := amqp091.Dial(config.Cfg.RabbitMQ.Address)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		s.rabbitmqConn = conn
	}

	return s.rabbitmqConn
}

func (s *ServiceProvider) RabbitMQChannel() *amqp091.Channel {
	if s.rabbitMqCh == nil {
		ch, err := s.RabbitMQConnection().Channel()
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		s.rabbitMqCh = ch
	}

	return s.rabbitMqCh
}
