package queues

import (
	"log"

	"github.com/alexis-aguirre/RabbitMQ-crash-course/storage/util"
	"github.com/streadway/amqp"
)

const (
	EXCHANGE_TYPE string = "topic"
)

var globalConfig util.Config

type queueManager struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func NewQueueManager(connectionString string) (queueManager, error) {
	globalConfig = util.GetConfig()
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return queueManager{}, err
	}
	return queueManager{connection: connection}, nil
}

func (qm *queueManager) Close() {
	qm.channel.Close()
	qm.connection.Close()
}

func (qm *queueManager) SetupQueues() {
	channel := createRabbitChannel(qm.connection)
	qm.channel = channel

	queueConfig := globalConfig.QueueConfig
	qm.declareExchange(globalConfig.QueueConfig.ExchangeName)
	qm.declareQueue(globalConfig.QueueConfig.QueueName, nil)
	qm.bindQueue(queueConfig.QueueName, queueConfig.RoutingKey, queueConfig.ExchangeName, nil)
}

func createRabbitChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()
	if err != nil {
		log.Fatal("Cannot create RabbitMQ channel: " + err.Error())
	}
	return channel
}

func (qm *queueManager) declareQueue(queueName string, args amqp.Table) {
	_, err := qm.channel.QueueDeclare(queueName, true, false, false, false, args)

	if err != nil {
		log.Fatal("Cannot declare queue " + queueName)
	}
	log.Println("Queue '" + queueName + "' created")
}

func (qm *queueManager) bindQueue(queue string, routingKey string, exchange string, args amqp.Table) {
	qm.channel.QueueBind(queue, routingKey, exchange, false, args)
}

func (qm *queueManager) declareExchange(exchangeName string) {
	err := qm.channel.ExchangeDeclare(exchangeName, EXCHANGE_TYPE, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exchange '" + exchangeName + "' created")
}
