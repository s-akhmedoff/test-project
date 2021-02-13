env_up:
	@docker-compose up -d
	@echo "Created, please be patient, starting service may take a few moments"

env_down:
	@docker-compose down

swag:
	~/go/bin/swag init -g api/api.go api/docs

api_up:
	go run cmd/main.go

generator_up:
	go run internal/generator/server.go

bot_up:
	go run internal/bot/subscriber.go

example_low:
	curl -d '{"number": 3, "priority": "low"}' -H 'Content-Type: application/json' localhost:8080/send

example_med:
	curl -d '{"number": 3, "priority": "medium"}' -H 'Content-Type: application/json' localhost:8080/send

example_high:
	curl -d '{"number": 3, "priority": "high"}' -H 'Content-Type: application/json' localhost:8080/send

post: example_low example_med example_high

proto:
	protoc -I=internal/generator/protos -I /usr/local/include --go_out=plugins=grpc:internal/generator/genproto internal/generator/protos/*.proto
	protoc -I=api/protos -I /usr/local/include --go_out=plugins=grpc:api/genproto  api/protos/*.proto