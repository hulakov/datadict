package main

import (
	"github.com/hulakov/datadict/pkg/logger"
	"github.com/hulakov/datadict/server/ports/rabbitmq"
	"github.com/rs/zerolog/log"
)

func main() {
	flags := ParseFlags()
	logger.Init(flags.Verbose)

	err := rabbitmq.Run(flags.RabbitMQHost)
	if err != nil {
		log.Error().Err(err).Msg("server failed")
	}

}
