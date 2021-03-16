package queues

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/alexis-aguirre/RabbitMQ-crash-course/dto"
	"github.com/streadway/amqp"
)

const (
	RETRY_COUNT_HEADER = "retry-count"

	CONTENT_TYPE_APPLICATION_JSON = "application/json"
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

		fmt.Println("Message Received: " + fmt.Sprint(obj))

		if !obj.Validate() {
			qm.moveToParkingLot(message)
		}

		message.Ack(false)
	}
}

func (qm *queueManager) moveToParkingLot(message amqp.Delivery) {
	log.Println("Moving message to parking lot: " + fmt.Sprint(string(message.Body)))
	queueConfig := globalConfig.QueueConfig
	var messageHeaders = amqp.Table{}
	if message.Headers == nil {
		messageHeaders[RETRY_COUNT_HEADER] = 1
	} else {
		retryCount := message.Headers[RETRY_COUNT_HEADER].(int32)
		messageHeaders[RETRY_COUNT_HEADER] = retryCount + 1
	}

	err := qm.channel.Publish(queueConfig.ExchangeName, queueConfig.RoutingKey+PARKING_SUFIX, false, false, amqp.Publishing{
		ContentType: CONTENT_TYPE_APPLICATION_JSON,
		Body:        message.Body,
		Headers:     messageHeaders,
	})

	if err != nil {
		log.Println("Error: " + err.Error())
	}

}
