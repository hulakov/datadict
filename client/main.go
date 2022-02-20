package main

import (
	"github.com/hulakov/datadict/pkg/logger"
	"github.com/hulakov/datadict/pkg/rabbitrpc"
	"github.com/rs/zerolog/log"
)

func processCommands(rabbitMQHost string, commands []string) {
	responses := make(chan []byte)
	requests := make(chan []byte)
	go func() {
		err := rabbitrpc.RunClient(rabbitMQHost, requests, responses)
		if err != nil {
			log.Error().Err(err).Msg("RabbitMQ client exited with error")
		}
	}()

	numRequests := 0
	for _, command := range commands {
		log.Debug().Str("command", string(command)).Msg("process command")
		request, err := commandToJson(command)
		if err != nil {
			log.Error().Str("command", command).Err(err).Msg("bad command")
			continue
		}
		requests <- request
		numRequests++
	}

	for numRequests > 0 {
		response := <-responses
		printResponse(response)
		numRequests--
	}

	log.Info().Msg("done")
}

func main() {
	flags := ParseFlags()
	logger.Init(flags.Verbose)
	processCommands(flags.RabbitMQHost, flags.Commands)
}
