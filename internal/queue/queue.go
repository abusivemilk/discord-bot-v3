package queue

import (
	"fmt"
	"github.com/VATUSA/discord-bot-v3/internal/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", config.QueueUser, config.QueuePassword, config.QueueHost, config.QueuePort)
}

func Connect() (*amqp.Connection, *amqp.Channel, amqp.Queue) {
	conn, err := amqp.Dial(ConnectionString())
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"discord_sync", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return conn, ch, q
}
