package notifications_adapter

import (
	"encoding/json"
	"errors"

	"github.com/GP-Hacks/auth/internal/config"
	"github.com/GP-Hacks/auth/internal/models"
	"github.com/GP-Hacks/auth/internal/utils/errs"
	"github.com/rabbitmq/amqp091-go"
)

func (a *NotificationsAdapter) SendMail(m *models.Mail) error {
	bodyBytes, err := json.Marshal(m)
	if err != nil {
		return errors.Join(errs.SomeError, err)
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
		return errors.Join(errs.SomeError, err)
	}

	return nil
}
