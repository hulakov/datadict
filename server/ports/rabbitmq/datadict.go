package rabbitmq

import (
	"github.com/hulakov/datadict/pkg/rabbitrpc"
)

func process(request []byte) []byte {
	s := "received: " + string(request)
	return ([]byte)(s)
}

func Run(rabbitMQHost string) error {
	return rabbitrpc.RunServer(rabbitMQHost, process)
}
