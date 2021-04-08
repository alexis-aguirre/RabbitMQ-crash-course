package queues

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexis-aguirre/RabbitMQ-crash-course/dto"
	"github.com/alexis-aguirre/RabbitMQ-crash-course/services/imageProcessingService"
	"github.com/alexis-aguirre/RabbitMQ-crash-course/util"
)

func (qm *queueManager) ListenOnQueue() {
	globalConfig := util.GetConfig()
	imageClient := imageProcessingService.NewImageProcessingClient(globalConfig.ServicesConfig.ImageProcessingUrl)
	queueConfig := globalConfig.QueueConfig
	log.Println("Listening on queue '" + queueConfig.QueueName + "'")

	messages, err := qm.channel.Consume(queueConfig.QueueName, "", false, false, false, false, nil)

	if err != nil {
		log.Fatal("Cannot consume from queue " + queueConfig.QueueName)
	}

	for message := range messages {
		obj := dto.ImageReport{}
		json.Unmarshal(message.Body, &obj)

		log.Println("Message Received: " + fmt.Sprint(obj))
		err = imageClient.ProcessPlate(obj)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Processed ", fmt.Sprint(obj))
		}

		message.Ack(false)
	}
}
