package bot

import (
	"github.com/VATUSA/discord-bot-v3/internal/queue"
	"github.com/bwmarrin/discordgo"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func QueueListen(s *discordgo.Session) {
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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	queue.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			ProcessMemberInGuilds(s, string(d.Body))
			d.Ack(false)
		}
		log.Printf("End of MQ loop")
	}()

	log.Print("Connected to RabbitMQ. Waiting for messages...")
	<-forever
	log.Print("Stopped processing RabbitMQ.")
}
