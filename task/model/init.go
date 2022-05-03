package model

import "github.com/streadway/amqp"

var MQ *amqp.Connection

func RabbitMQConnect(connect string) {
	conn, err := amqp.Dial(connect)
	if err != nil {
		panic(err)
	}
	MQ = conn
}
