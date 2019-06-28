package interfaces

import (
	"encoding/json"
	"github.com/adriancarayol/go-event/config/rabbit"
	"github.com/adriancarayol/go-event/pkg/dao"
	"github.com/streadway/amqp"
	"log"
)

type eventBusRabbitMQ struct {
	amqp *amqp.Connection
}

func NewEventBus() *eventBusRabbitMQ {
	return &eventBusRabbitMQ{
		amqp: rabbit.Get(),
	}
}

func (b *eventBusRabbitMQ) Publish(event dao.Event) error {
	amqpConnection := rabbit.Get()
	ch, err := amqpConnection.Channel()

	if err != nil {
		log.Printf("Error creating rabbitmq channel: %s", err)
		return err
	}

	payload, err := json.Marshal(event)

	if err != nil {
		log.Printf("Error marshal event: %s", err)
		return err
	}

	err = ch.Publish("", "user_events", false, false, amqp.Publishing{DeliveryMode: amqp.Persistent, ContentType: "text/plain", Body: payload})

	if err != nil {
		log.Printf("Error publishing message in rabbitmq: %s", err)
		return err
	}

	return nil
}
