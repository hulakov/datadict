package rabbitrpc

import (
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/streadway/amqp"
)

// The implementation is based on https://www.rabbitmq.com/tutorials/tutorial-six-go.html
func RunServer(host string, process func(request []byte) []byte) error {
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
		ServerQueueName, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %w", err)
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return fmt.Errorf("failed to set QoS: %w", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(NumWorkers)

	worker := func() {
		for d := range msgs {
			log.Debug().
				Str("body", string(d.Body)).
				Str("correlation_id", d.CorrelationId).
				Msg("received request")

			response := process(d.Body)
			log.Debug().
				Str("body", string(response)).
				Str("correlation_id", d.CorrelationId).
				Msg("send response")

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   ContentType,
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
			if err != nil {
				log.Debug().
					Str("body", string(d.Body)).
					Str("correlation_id", d.CorrelationId).
					Err(err).
					Msg("failed to publish a message")
			}

			d.Ack(false)
		}

		wg.Done()
	}

	for i := 0; i < NumWorkers; i++ {
		go worker()
	}
	wg.Wait()

	return nil
}
