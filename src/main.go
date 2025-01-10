package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	tele "gopkg.in/telebot.v4"
)

type Data struct {
}

var TestData string

func main() {
	pref := tele.Settings{
		Token:  "7434140671:AAGineKwMZ-T6_I0vA92qbcrC0K8A9R7YdU",
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

	b.Handle("/test", func(c tele.Context) error {
		GetDataFromApi()
		return c.Send(TestData)
	})

	b.Handle(&btnPrev, func(c tele.Context) error {
		return c.Send("Choose cryptocurrency:", menu)
	})
	b.Start()
}

func GetDataFromApi() {
	url := "https://api.coingecko.com/api/v3/coins/solana?x_cg_demo_api_key=CG-pXdsuqQf4duHNu3i1L2JetMp"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	TestData = string(body)
}
