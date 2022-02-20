package rabbitrpc

const (
	DefaultHost     = "amqp://guest:guest@localhost:5672/"
	ContentType     = "application/json"
	ServerQueueName = "rpc_queue"
	NumWorkers      = 10
)
