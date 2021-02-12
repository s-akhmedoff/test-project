package handler

import (
	"context"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"test-project/internal/generator/genproto"
	"test-project/utils"
)

type (
	message struct {
		text []byte
	}

	messages []message

	Server struct {}
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters = 10
)

func randStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func generateMessages(number uint32) messages {
	messageBunch := make(messages, number)
	for i := range messageBunch {
		messageBunch[i].text = randStringBytes(letters)
	}
	return messageBunch
}

//GenerateAndSend - ...
func (s *Server) GenerateAndSend(ctx context.Context, message *genproto.GSRequest) (*genproto.GSResponse, error) {

	messageBunch := generateMessages(message.Number)
	log.Print(messageBunch)
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	ch.ExchangeDeclare("tg", "topic", true, false, false, false, nil)

	log.Print(message.Priority)
	for i := range messageBunch {
		body := messageBunch[i].text
		err := ch.Publish(
			"tg",                   // exchange
			message.Priority, // routing key
			false,                  // mandatory
			false,                  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
			})
		utils.FailOnError(err, "Failed to publish a message")

		log.Printf("Message '%s' is sent", body)
	}
	return &genproto.GSResponse{Response: "OK"}, nil
}