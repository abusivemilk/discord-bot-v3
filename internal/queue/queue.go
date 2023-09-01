package queue

import (
	"fmt"
	"github.com/VATUSA/discord-bot-v3/internal/config"
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
