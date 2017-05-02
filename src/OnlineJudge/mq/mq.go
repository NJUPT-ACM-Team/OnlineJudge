package mq

import (
	"github.com/streadway/amqp"

	"fmt"
)

const (
	LJQueueName = "local"
	VJQueueName = "virtual"
	MJQueueName = "manual"
)

type MQConfig struct {
	Username string
	Password string
	Address  string
}

func (this *MQConfig) GetDSN() string {
	dsn := fmt.Sprintf("amqp://%s:%s@%s/",
		this.Username, this.Password, this.Address)
	return dsn
}

type MQ struct {
	DSN  string
	conn *amqp.Connection
	ch   *amqp.Channel
	ljq  amqp.Queue
	vjq  amqp.Queue
	mjq  amqp.Queue
}

var DSN string

func Init(cfg *MQConfig) {
	DSN = cfg.GetDSN()
}

func New() *MQ {
	return &MQ{
		DSN: DSN,
	}
}

func (this *MQ) Reset() {
	*this = MQ{
		DSN: DSN,
	}
}

func (this *MQ) Connect() error {
	var err error
	this.conn, err = amqp.Dial(this.DSN)
	if err != nil {
		return err
	}
	this.ch, err = this.conn.Channel()
	if err != nil {
		this.ch = nil
		return err
	}
	return nil
}

func (this *MQ) Disconnect() {
	this.conn.Close()
}

func (this *MQ) Declare(name string) (amqp.Queue, error) {
	return this.ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (this *MQ) DeclareLJ() error {
	var err error
	this.ljq, err = this.Declare(LJQueueName)
	if err != nil {
		return err
	}
	return nil
}

func (this *MQ) DeclareVJ() error {
	var err error
	this.vjq, err = this.Declare(VJQueueName)
	if err != nil {
		return err
	}
	return nil
}

func (this *MQ) DeclareMJ() error {
	var err error
	this.mjq, err = this.Declare(MJQueueName)
	if err != nil {
		return err
	}
	return nil
}

func (this *MQ) Publish(body []byte, qname string) error {
	err := this.ch.Publish(
		"",    // exchange
		qname, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	return err
}

func (this *MQ) PublishLJ(body []byte) error {
	return this.Publish(body, this.ljq.Name)
}

func (this *MQ) PublishVJ(body []byte) error {
	return this.Publish(body, this.vjq.Name)
}

func (this *MQ) PublishMJ(body []byte) error {
	return this.Publish(body, this.mjq.Name)
}

func (this *MQ) Receiver(qname string, fn func([]byte)) error {
	msgs, err := this.ch.Consume(
		qname, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}
	go func() {
		for d := range msgs {
			fn(d.Body)
		}
	}()
	return nil
}

func (this *MQ) LJReceiver(fn func([]byte)) error {
	return this.Receiver(this.ljq.Name, fn)
}

func (this *MQ) VJReceiver(fn func([]byte)) error {
	return this.Receiver(this.vjq.Name, fn)
}

func (this *MQ) MJReceiver(fn func([]byte)) error {
	return this.Receiver(this.mjq.Name, fn)
}
