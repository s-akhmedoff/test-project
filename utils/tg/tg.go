package tg

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"test-project/utils"
)

func InitBot() (bot *tgbotapi.BotAPI) {
	bot, err := tgbotapi.NewBotAPI("1672264352:AAGznTbLlxmLu6xJxYjzPEjU_0raes7gfWE")
	utils.FailOnError(err, "failed to connect to bot")

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return
}
