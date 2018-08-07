package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// conn, err := amqp.Dial("amqp://huangjia:huangjia@puhui@10.10.108.142:5672/backup")
	conn, err := amqp.Dial("amqp://rabbitmq:0TFXakhABBTonR7v@rabbitmq-cluster-0.rabbitmq-cluster.mysql-operator.svc.cluster.local:5672/backup")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
