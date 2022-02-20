# datadict

datadict is a client-server application which provides key-value in-memory storage which preserves items order.
The communication between client and server is established via RabbitMQ.


## Prerequrements
- golang 1.16.6+
- docker 20.10.10+
## How to run


1. Run RabbitMQ:
    ```
    make rabbitmq.run
    ```
    The command uses docker (i.e. it's expected that docker is up and running)
2. Run server via the command:
    ```
    go run ./server
    ```
    In order to run server in verbose mode use `--verbose` key:
    ```
    go run ./server --verbose
    ```
    By default server is connected to `amqp://guest:guest@localhost:5672/` but it's possible to override it:
    ```
    go run ./server -rabbitmq-host amqp://user:pass@rabbitsrv:5672/
    ```
3. At this point the system is ready to handle requests. The client CLI looks like this:
    ```
    go run ./client [--verbose] [-rabbitmq-host <host>] cmd1 [... cmdN]
    ```
    Where cmdN is one of:
    - `add_item:<key>:<value>`
    - `remove_item:<key>`
    - `get_item:<key>`
    - `get_all_items`

    For instance:
    ```
    go run ./client add_item:a,1 get_item:a add_item:b,2 add_item:d,3 add_item:e,4 add_item:c,5 remove_item:d get_all_items
    ```

## Implementation overview
The server creates queue for the requests. Every client creates another exclusive queue for the responses.
Each request is supplied with name of the queue for response (needed for client matching) and correlation ID (needed for request-response matching).

The data is stored in memory (i.e. everything is lost after server restart).
For storage and ordering used double-linked list. The indexing is implemented via Go map (hash table). The storage supports concurent access (uses read-write mutex).

The server processes requests in parallel (there are 10 workers by default).
The client supports parallel message sending and receiving. The implementation is done via goroutines and channels.

The application implements proper error handling.

## Project structure
- `server` - server sources. The implementation is done according to the ports and adapters architecture
    - `models` - domain data models
    - `ports` - server entry points
        - `rabbitmq` - RabbitMQ port
    - `services` - bussines logic implementation (since there is no complex logic so far there are only database access and logging)
    - `storage` - repositories declarations and implementationts
- `pkg` - packages shared between server and client
    - `datamsg` - JSON serializer/deserializer for client/server communication
    - `rabbitrpc` - implementation of RPC on top of RabbitMQ
    - `logger` - logger configuration
- `client` - client sources. The client is implemented in accordance to KISS principle



