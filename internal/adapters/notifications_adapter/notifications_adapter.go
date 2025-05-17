package notifications_adapter

import "github.com/rabbitmq/amqp091-go"

type NotificationsAdapter struct {
	ch *amqp091.Channel
}

func NewNotificationsAdapter(ch *amqp091.Channel) *NotificationsAdapter {
	return &NotificationsAdapter{
		ch: ch,
	}
}
