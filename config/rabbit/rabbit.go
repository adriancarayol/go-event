package rabbit

import (
	"github.com/streadway/amqp"
	"log"
)

var amqpConnection *amqp.Connection

func configureQueues() {
	ch, err := amqpConnection.Channel()
	if err != nil {
		log.Fatalf("Error creating channel: %s", err)
	}

	_, err = ch.QueueDeclare("user_events", true, true, false, false, nil)

	if err != nil {
		log.Fatalf("Error creating queue: %s", err)
	}

	if err := ch.Close(); err != nil {
		log.Fatalf("Error closing rabbitmq channel: %s", err)
	}
}

func Init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalf("Error connecting with RabbitMQ: %s", err)
	}

	amqpConnection = conn

	configureQueues()
}

func Get() *amqp.Connection {
	return amqpConnection
}
