package rabbitrpc

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

// The implementation is based on https://www.rabbitmq.com/tutorials/tutorial-six-go.html
func RunClient(host string, requests <-chan []byte, responses chan<- []byte) error {
	conn, err := amqp.Dial(host)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to open a queue: %w", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	nextMsgID := 1

	for {
		select {
		case request, ok := <-requests:
			if !ok {
				return nil
			}
			correlationId := fmt.Sprintf("msg%d", nextMsgID)
			nextMsgID++
			err = ch.Publish(
				"",              // exchange
				ServerQueueName, // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType:   ContentType,
					CorrelationId: correlationId,
					ReplyTo:       q.Name,
					Body:          request,
				})
			if err != nil {
				return fmt.Errorf("failed to to publish a message: %w", err)
			}
			log.Debug().
				Str("body", string(request)).
				Str("correlation_id", correlationId).
				Msg("sent request")
		case msg, ok := <-msgs:
			if !ok {
				return nil
			}
			log.Debug().
				Str("body", string(msg.Body)).
				Str("correlation_id", msg.CorrelationId).
				Msg("received response")
			responses <- msg.Body
		}
	}
}
