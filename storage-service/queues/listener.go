package queues

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexis-aguirre/RabbitMQ-crash-course/storage/dto"
)

func (qm *queueManager) ListenOnQueue() {
	queueConfig := globalConfig.QueueConfig
	log.Println("Listening on queue '" + queueConfig.QueueName + "'")

	messages, err := qm.channel.Consume(queueConfig.QueueName, "", false, false, false, false, nil)

	if err != nil {
		log.Fatal("Cannot consume from queue " + queueConfig.QueueName)
	}

	for message := range messages {
		obj := dto.ImageReport{}
		json.Unmarshal(message.Body, &obj)

		fmt.Println("Message Stored: " + fmt.Sprint(obj))

		message.Ack(false)

	}
}
