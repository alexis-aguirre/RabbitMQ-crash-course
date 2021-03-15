package main

import (
	"fmt"
	"log"

	"github.com/alexis-aguirre/RabbitMQ-crash-course/queues"
	"github.com/alexis-aguirre/RabbitMQ-crash-course/util"
)

func main() {
	log.Println("Starting rabbitMQ consumer")

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to read config")
	}

	rabbitConnectionString := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RabbitConfig.RabbitUser, config.RabbitConfig.RabbitPassword, config.RabbitConfig.HOST, config.RabbitConfig.PORT)
	queueManager, err := queues.NewQueueManager(rabbitConnectionString)

	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}
	queueManager.SetupQueues()
	defer queueManager.Close()

	forever := make(chan bool)

	go func() {
		queueManager.ListenOnQueue()
	}()

	<-forever

}
