package web

import (
	"context"
	"github.com/VATUSA/discord-bot-v3/internal/queue"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func SendToQueue(id string) {
	conn, err := amqp.Dial(queue.ConnectionString())
	queue.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	queue.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"discord_sync", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	queue.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(id),
	})
	queue.FailOnError(err, "Failed to publish message")
}
