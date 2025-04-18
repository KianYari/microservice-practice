package main

import (
	"fmt"
	"log"

	"github.com/kianyari/microservice-practice/notification-service/config"
	"github.com/streadway/amqp"
)

func main() {
	cfg := config.LoadConfig()
	rabbitMQURL := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RabbitMQUser,
		cfg.RabbitMQPassword,
		cfg.RabbitMQHost,
		cfg.RabbitMQPort,
	)
	rabbitMQConn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitMQConn.Close()

	ch, err := rabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"task_deadline",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare an exchange: %v", err)
	}
	queue, err := ch.QueueDeclare(
		"task_deadline_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		queue.Name,
		"task_deadline",
		"task_deadline",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to bind the queue: %v", err)
	}
	msgs, err := ch.Consume(
		queue.Name,
		"task_deadline_consumer",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}

	err = ch.ExchangeDeclare(
		"task_created",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare an exchange: %v", err)
	}
	cqueue, err := ch.QueueDeclare(
		"task_created_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare a queue: %v", err)
	}
	err = ch.QueueBind(
		cqueue.Name,
		"task_created",
		"task_created",
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to bind the queue: %v", err)
	}
	cmsgs, err := ch.Consume(
		cqueue.Name,
		"task_created_consumer",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer: %v", err)
	}
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)
			// Here you can add code to handle the notification, e.g., send an email or push notification
		}
	}()
	go func() {
		for cmsg := range cmsgs {
			log.Printf("Received a message: %s", cmsg.Body)
			// Here you can add code to handle the notification, e.g., send an email or push notification
		}
	}()
	log.Println("Waiting for messages. To exit press CTRL+C")
	forever := make(chan bool)
	<-forever
	log.Println("Exiting...")
}
