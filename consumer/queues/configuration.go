package queues

import (
	"log"

	"github.com/alexis-aguirre/RabbitMQ-crash-course/util"
	"github.com/streadway/amqp"
)

const (
	EXCHANGE_TYPE           string = "topic"
	PARKING_SUFIX           string = ".parking"
	DEAD_LETTER_QUEUE_SUFIX string = ".dead-letter-queue"

	NO_TTL int = 0
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
	//Dead letter queue
	qm.declareQueue(globalConfig.QueueConfig.QueueName+DEAD_LETTER_QUEUE_SUFIX, nil)

	//Main queue
	mainQueueParameters := configureDeadLetterQueue(globalConfig.QueueConfig.ExchangeName, globalConfig.QueueConfig.RoutingKey+DEAD_LETTER_QUEUE_SUFIX, NO_TTL)
	qm.declareQueue(globalConfig.QueueConfig.QueueName, mainQueueParameters)

	//Parking lot queue
	parkingLotQueueParameters := configureDeadLetterQueue(globalConfig.QueueConfig.ExchangeName, globalConfig.QueueConfig.RoutingKey, 10000)
	qm.declareQueue(globalConfig.QueueConfig.QueueName+PARKING_SUFIX, parkingLotQueueParameters)

	qm.bindQueue(queueConfig.QueueName, queueConfig.RoutingKey, queueConfig.ExchangeName, nil)                                                 //Main queue
	qm.bindQueue(queueConfig.QueueName+PARKING_SUFIX, queueConfig.RoutingKey+PARKING_SUFIX, queueConfig.ExchangeName, nil)                     //Parking lot queue
	qm.bindQueue(queueConfig.QueueName+DEAD_LETTER_QUEUE_SUFIX, queueConfig.RoutingKey+DEAD_LETTER_QUEUE_SUFIX, queueConfig.ExchangeName, nil) //Dead letter queue
}

func configureDeadLetterQueue(exchange string, routingKey string, ttl int) amqp.Table {
	parameters := amqp.Table{}
	parameters["x-dead-letter-exchange"] = exchange
	parameters["x-dead-letter-routing-key"] = routingKey
	if ttl != NO_TTL {
		parameters["x-message-ttl"] = ttl
	}

	return parameters
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
		log.Println(err)
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
