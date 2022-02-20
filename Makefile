rabbitmq.start:
	docker run -d --hostname my-rabbit --name some-rabbit rabbitmq:3

rabbitmq.stop:
	docker stop some-rabbit
	docker rm some-rabbit

rabbitmq.run:
	docker run -it --rm --hostname my-rabbit --name some-rabbit rabbitmq:3

server.run:
	go run ./server/main.go

client.run:
	go run ./client/main.go