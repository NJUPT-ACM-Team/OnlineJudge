package mq

import (
	"github.com/streadway/amqp"
)

var DSN string

type MQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	ljq  amqp.Queue
	vjq  amqp.Queue
}

func Init() {
	DSN = "amqp://guest:guest@localhost:5672/"
}

func Declare() error {
	return nil
}
