package notifications_adapter

import (
	"encoding/json"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/services"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

func (a *NotificationsAdapter) SendMail(m *models.Mail) error {
	bodyBytes, err := json.Marshal(m)
	if err != nil {
		log.Error().Msg(err.Error())

		return services.InternalServer
	}

	err = a.ch.Publish(
		"",
		config.Cfg.RabbitMQ.EmailQueue,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        bodyBytes,
		},
	)
	if err != nil {
		log.Error().Msg(err.Error())

		return services.InternalServer
	}

	return nil
}
