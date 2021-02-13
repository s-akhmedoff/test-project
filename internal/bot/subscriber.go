package main

import (
	"log"
	"test-project/utils"
	"test-project/utils/tg"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/streadway/amqp"
)

var priorities = []string{"low", "medium", "high"}

func main() {

	bot := tg.InitBot()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	ch.ExchangeDeclare("tg", "topic", true, false, false, false, nil)

	_ = ch.Qos(1, 0, false)

	qs := make(map[string]<-chan amqp.Delivery)

	for _, p := range priorities {
		q, err := ch.QueueDeclare(
			"",    // name
			false, // durable
			false, // delete when unused
			true,  // exclusive
			false, // no-wait
			nil,   // arguments
		)
		utils.FailOnError(err, "Failed to declare a queue")

		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "tg", p)
		err = ch.QueueBind(
			q.Name, // queue name
			p,      // routing key
			"tg",   // exchange
			false,
			nil)
		utils.FailOnError(err, "Failed to bind a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto ack
			false,  // exclusive
			false,  // no local
			false,  // no wait
			nil,    // args
		)
		utils.FailOnError(err, "Failed to register a consumer")

		qs[p] = msgs

	}

	forever := make(chan bool)

	var current *amqp.Delivery
	go func() {
		for {
			time.Sleep(time.Second)

			select {
			case msg := <-qs["high"]:
				current = &msg
			default:
				current = nil
			}
			if current == nil {
				select {
				case msg := <-qs["medium"]:
					current = &msg
				default:
					current = nil
				}
			}
			if current == nil {
				select {
				case msg := <-qs["low"]:
					current = &msg
				default:
					current = nil
					break
				}
			}
			if current != nil {
				log.Print(current)
				bot.Send(tgbotapi.NewMessage(int64(-1001332343159), string(current.Body)))
			}
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
