package main

import (
	"fmt"
	"os"

	"github.com/mymmrac/telego"

	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := ""

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		chatId := tu.ID(update.Message.Chat.ID)
		_, _ = bot.SendSticker(
			tu.Sticker(
				chatId,
				tu.FileFromID("CAACAgIAAxkBAAENbgRndo9rt-tWiih7QglKDk4jd9i9PQACAwEAAladvQoC5dF4h-X6TzYE"),
			),
		)
		message := tu.Message(
			chatId,
			"Hello!",
		)

		bot.SendMessage(message)

	}, th.CommandEqual("start"))
	bh.Start()
}
