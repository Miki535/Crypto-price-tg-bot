package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 1 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	//var lastmessageText string

	var (
		selector = &tele.ReplyMarkup{}

		menu = &tele.ReplyMarkup{ResizeKeyboard: true}

		btnBtc  = menu.Text("Bitcoin")
		btnUsdt = menu.Text("USDT")
		btnEth  = menu.Text("Ethereum")
		btnSol  = menu.Text("Solana")
		btnTon  = menu.Text("Toncoin")

		//Inline buttons
		btnPrev = selector.Data("₿ Cryptocurrency", "crypto")
	)

	menu.Reply(
		menu.Row(btnBtc),
		menu.Row(btnUsdt),
		menu.Row(btnEth),
		menu.Row(btnSol),
		menu.Row(btnTon),
	)

	selector.Inline(
		selector.Row(btnPrev),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! \nHere you can find the latest cryptocurrency rates for Bitcoin, Ethereum, Solana, USDT, and TON", selector)
	})

	b.Handle(&btnPrev, func(c tele.Context) error {
		return c.Send("Choose cryptocurrency:", menu)
	})
	b.Start()
}
