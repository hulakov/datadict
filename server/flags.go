package main

import (
	"flag"

	"github.com/hulakov/datadict/pkg/rabbitrpc"
)

type Flags struct {
	RabbitMQHost string
	Verbose      bool
}

func ParseFlags() Flags {
	var flags Flags
	flag.StringVar(&flags.RabbitMQHost, "rabbitmq-host", rabbitrpc.DefaultHost, "RabbitMQ host")
	flag.BoolVar(&flags.Verbose, "verbose", false, "Print debug info")
	flag.Parse()
	return flags
}
