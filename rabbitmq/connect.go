package rabbitmq

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (rabbitmq *RabbitMQ) failOnError(err error, msg string) {
	if err != nil {
		log.Error(err)
	}
}

func (rabbitmq *RabbitMQ) GetQueueName() string {
	return rabbitmq.exchangeName
}

// CreateCon create connection
func (rabbitmq *RabbitMQ) CreateCon(connectionString string) *amqp.Connection {
	conn, err := amqp.Dial(connectionString)
	rabbitmq.failOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

// CreateChannel create channel
func (rabbitmq *RabbitMQ) CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		rabbitmq.failOnError(err, "Failed to open a channel")
		return nil
	}
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		rabbitmq.failOnError(err, "Failed to set QoS")
		return nil
	}
	return ch
}

// CreateQueue create queue
func (rabbitmq *RabbitMQ) CreateQueue(rbmq *amqp.Channel, queuename string) amqp.Queue {
	q, err := rbmq.QueueDeclare(
		queuename, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		// arg,       // arguments
		nil,
	)
	if err != nil {
		rabbitmq.failOnError(err, "Failed to declare a queue")
	}
	return q
}

func (rabbitmq *RabbitMQ) CreateExchange(ch *amqp.Channel, name string) {
	err := ch.ExchangeDeclare(
		name,     // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		rabbitmq.failOnError(err, "Failed to declare a queue")
	}
}

type RabbitMQ struct {
	queue        amqp.Queue
	rbCon        *amqp.Connection
	ch           *amqp.Channel
	exchangeName string
}

func (rabbitmq *RabbitMQ) Publish(msg interface{}) error {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = rabbitmq.ch.Publish(
		rabbitmq.exchangeName, // exchange
		"",                    // routing key
		false,                 // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         bytes,
		})
	return err
}

func New(conn, queueName, queueNameRetry string) *RabbitMQ {
	rabbitmq := &RabbitMQ{}
	rabbitCon := rabbitmq.CreateCon(conn)
	rabbitmq.ch = rabbitmq.CreateChannel(rabbitCon)
	rabbitmq.queue = rabbitmq.CreateQueue(rabbitmq.ch, queueName)
	rabbitmq.exchangeName = queueName
	return rabbitmq
}
