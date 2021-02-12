env_up:
	docker run -d --name rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:3-management

swag:
	~/go/bin/swag init -g api/api.go api/docs

api_up:
	go run cmd/main.go

generator_up:
	go run internal/generator/server.go

bot_up:
	go run internal/bot/subscriber.go

example_low:
	curl -d '{"number": 10, "priority": "low"}' -H 'Content-Type: application/json' localhost:8080/send

example_med:
	curl -d '{"number": 10, "priority": "medium"}' -H 'Content-Type: application/json' localhost:8080/send

example_high:
	curl -d '{"number": 10, "priority": "high"}' -H 'Content-Type: application/json' localhost:8080/send

post: example_low example_med example_high

proto:
	protoc -I=internal/generator/protos -I /usr/local/include --go_out=plugins=grpc:internal/generator/genproto internal/generator/protos/*.proto
	protoc -I=api/protos -I /usr/local/include --go_out=plugins=grpc:api/genproto  api/protos/*.proto